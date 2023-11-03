package server

import (
	"fmt"

	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"github.com/boxcolli/go-transistor/core"
	"github.com/boxcolli/go-transistor/emitter/basicemitter"
	"github.com/boxcolli/go-transistor/io/reader/grpcreader"
	"github.com/boxcolli/go-transistor/io/writer/grpcwriter"
	"github.com/boxcolli/go-transistor/server"
	"github.com/boxcolli/go-transistor/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInternal = status.Error(codes.Internal, "failed to serve request.")
)

type grpcServer struct {
	pb.UnimplementedTransistorServiceServer

	eqs int

	c core.Core
}

func NewTransistorServer(c core.Core, emitterQueueSiz int) server.Server {
	return &grpcServer{
		c: c,
		eqs: emitterQueueSiz,
	}
}

// Publish implements pb.TransistorServiceServer.
func (s *grpcServer) Publish(stream pb.TransistorService_PublishServer) error {
	fmt.Printf("Begin\n")
	r := grpcreader.NewGrpcServerStream(stream)
	fmt.Printf("New grpc server stream: %v\n", r)
	err := s.c.Collect(r)
	return err
}

// Subscribe implements pb.TransistorServiceServer.
func (s *grpcServer) Subscribe(stream pb.TransistorService_SubscribeServer) error {
	e := basicemitter.NewBasicEmitter(s.eqs)
	ch := make(chan error)

	// Receive at least one change
	{
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		cg := new(types.Change)
		cg.Unmarshal(req.GetChange())
		s.c.Apply(e, cg)
	}

	// Listen change
	go func() {
		req, err := stream.Recv()
		if err != nil {
			ch <- err
			return
		}

		cg := new(types.Change)
		cg.Unmarshal(req.GetChange())
		s.c.Apply(e, cg)
	} ()

	// Send message
	go func() {
		w := grpcwriter.NewGrpcWriter(stream)
		err := e.Work(w)
		ch <- err
	} ()

	err := <- ch
	s.c.Delete(e)
	return err
}

// Command implements pb.TransistorServiceServer.
func (s *grpcServer) Command(req *pb.CommandRequest, stream pb.TransistorService_CommandServer) error {
	// Command
	ch := s.c.Command(req.GetArgs())

	// Check at least one response
	success := false

	// Receive lines
	for {
		line, ok := <- ch
		if !ok {
			break	// finished
		}

		success = true

		// Send response line
		err := stream.Send(&pb.CommandResponse{ Line: line })
		if err != nil {
			return err
		}
	}

	if !success {
		return ErrInternal
	}
	return nil
}
