package reader

import "google.golang.org/grpc"

type clientStreamReader struct {
	
}

func NewGrpcClientStream(stream grpc.ClientStream) StreamReader {
	return &clientStreamReader{}
}

type serverStreamReader struct {

}

func NewGrpcServerStream(stream grpc.ServerStream) StreamReader {
	return &serverStreamReader{}	
}
