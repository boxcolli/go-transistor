package server

import (
	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/collector"
	"github.com/boxcolli/go-transistor/emitter/basicemitter"
	"github.com/boxcolli/go-transistor/io/reader/grpcreader"
	"github.com/boxcolli/go-transistor/io/writer/grpcwriter"
	"github.com/boxcolli/go-transistor/server"
	"github.com/boxcolli/go-transistor/types"
)

type grpcServer struct {
	pb.UnimplementedTransistorServiceServer

	qsiz int

	c collector.Collector
	b base.Base
}

// Publish implements pb.TransistorServiceServer.
func (s *grpcServer) Publish(stream pb.TransistorService_PublishServer) error {
	sr := grpcreader.NewGrpcServerStream(stream)
	err := s.c.Work(sr) // Block
	return err
}

// Subscribe implements pb.TransistorServiceServer.
func (s *grpcServer) Subscribe(stream pb.TransistorService_SubscribeServer) error {
	e := basicemitter.NewBasicEmitter(s.qsiz)
	ch := make(chan error)

	// Listen change
	go func() {
		req, err := stream.Recv()
		if err != nil {
			ch <- err
			return
		}

		cg := new(types.Change)
		cg.Unmarshal(req.GetChange())
		s.b.Apply(e, cg)
	} ()

	// Send message
	go func() {
		w := grpcwriter.NewGrpcWriter(stream)
		err := e.Work(w)
		ch <- err
	} ()

	err := <- ch
	s.b.Delete(e)
	return err
}

// Command implements pb.TransistorServiceServer.
func (s *grpcServer) Command(*pb.CommandRequest, pb.TransistorService_CommandServer) error {
	panic("unimplemented")
}

func NewTransistorServer() server.Server {
	return &grpcServer{}
}
