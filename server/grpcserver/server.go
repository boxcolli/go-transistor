package grpcserver

import (
	"sync"

	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"github.com/boxcolli/go-transistor/core"
	"github.com/boxcolli/go-transistor/io/reader/grpcreader"
	"github.com/boxcolli/go-transistor/io/writer/grpcwriter"
	"github.com/boxcolli/go-transistor/server"
	"github.com/boxcolli/go-transistor/types"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInternal = status.Error(codes.Internal, "failed to serve request.")
)

type grpcServer struct {
	core core.Core

	pb.UnimplementedTransistorServiceServer
}

func NewGrpcServer(core core.Core) server.Server {
	zerolog.New(zerolog.NewConsoleWriter())
	return &grpcServer{
		core: core,
	}
}

func (s *grpcServer) Publish(stream pb.TransistorService_PublishServer) error {
	r := grpcreader.NewGrpcServerStream(stream)
	err := s.core.Collect(r)
	return err
}

func (s *grpcServer) Subscribe(stream pb.TransistorService_SubscribeServer) error {
	w := grpcwriter.NewGrpcWriter(stream)
	{
		go s.core.Emit(w)
		defer s.core.Stop(w)
	}

	// Receive at least one change
	{
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		cg := new(types.Change)
		cg.Unmarshal(req.GetChange())
		s.core.Apply(w, cg)
	}

	// Listen change
	ch := make(chan error, 2)
	go func() {
		for {
			req, err := stream.Recv()
			if err != nil {
				ch <- err
				return
			}
	
			cg := new(types.Change)
			cg.Unmarshal(req.GetChange())
			s.core.Apply(w, cg)
		}
	} ()

	// Send message
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		w := grpcwriter.NewGrpcWriter(stream)
		err := s.core.Emit(w)	// block
		ch <- err
	} ()

	<- ch	// block
	wg.Wait()
	return nil
}

func (s *grpcServer) Command(req *pb.CommandRequest, stream pb.TransistorService_CommandServer) error {
	// Command
	ch, err := s.core.Command(stream.Context(), req.GetArgs())
	if err != nil {
		return ErrInternal
	}

	// Receive lines
	for {
		line, ok := <- ch
		if !ok { break }	// finished

		// Send response line
		err := stream.Send(&pb.CommandResponse{ Line: line })
		if err != nil { return err }
	}

	return nil
}
