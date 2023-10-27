package server

import (
	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/collector"
	"github.com/boxcolli/go-transistor/io/reader/grpcreader"
)

type transistorServer struct {
	pb.UnimplementedTransistorServiceServer

	c collector.Collector
	b base.Base
}

// Publish implements pb.TransistorServiceServer.
func (s *transistorServer) Publish(server pb.TransistorService_PublishServer) error {
	sr := grpcreader.NewGrpcServerStream(server)

	err := s.c.Work(sr) // Block


}

// Subscribe implements pb.TransistorServiceServer.
func (s *transistorServer) Subscribe(server pb.TransistorService_SubscribeServer) error {
	panic("unimplemented")

	go func() {
		// err := Emitter.Work()

		// err handle
	} ()

	for {
		req, err := server.Recv()

		cg := req.GetChange()

		s.b.Apply(e, cg)
	}
}

// Command implements pb.TransistorServiceServer.
func (*transistorServer) Command(*pb.CommandRequest, pb.TransistorService_CommandServer) error {
	panic("unimplemented")
}

func NewTransistorServer() pb.TransistorServiceServer {
	return &transistorServer{}
}
