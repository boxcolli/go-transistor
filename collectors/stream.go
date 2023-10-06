package collectors

import "google.golang.org/grpc"

type StreamAdapter interface {
	// Recv()
}

type clientStreamAdapter struct {

}

func NewClientStreamAdapter(stream grpc.ClientStream) StreamAdapter {
	return &clientStreamAdapter{}
}

type serverStreamAdapter struct {

}

func NewServerStreamAdapter(stream grpc.ServerStream) StreamAdapter {
	return &serverStreamAdapter{}	
}
