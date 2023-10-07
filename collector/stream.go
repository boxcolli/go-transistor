package collector

import "google.golang.org/grpc"

// Stream is an adapter of any streams like grpc.Stream.
type Stream interface {
	// Recv()
}

type clientStream struct {
	
}

func NewGrpcClientStream(stream grpc.ClientStream) Stream {
	return &clientStream{}
}

type serverStream struct {

}

func NewGrpcServerStream(stream grpc.ServerStream) Stream {
	return &serverStream{}	
}
