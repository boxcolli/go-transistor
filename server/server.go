package server

import pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"

type transistorServer struct {
	pb.UnimplementedTransistorServiceServer
}

// Command implements pb.TransistorServiceServer.
func (*transistorServer) Command(*pb.CommandRequest, pb.TransistorService_CommandServer) error {
	panic("unimplemented")
}

// Publish implements pb.TransistorServiceServer.
func (*transistorServer) Publish(pb.TransistorService_PublishServer) error {
	panic("unimplemented")
}

// Subscribe implements pb.TransistorServiceServer.
func (*transistorServer) Subscribe(pb.TransistorService_SubscribeServer) error {
	panic("unimplemented")
}

func NewTransistorServer() pb.TransistorServiceServer {
	return &transistorServer{}
}
