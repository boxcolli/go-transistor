package grpc

import (
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
	"google.golang.org/grpc"
)

type clientStreamReader struct {
}

// Read implements StreamReader.
func (*clientStreamReader) Read() (*types.Message, error) {
	panic("unimplemented")
}

func NewGrpcClientStream(stream grpc.ClientStream) io.StreamReader {
	return &clientStreamReader{}
}

type serverStreamReader struct {
}

// Read implements StreamReader.
func (*serverStreamReader) Read() (*types.Message, error) {
	panic("unimplemented")
}

func NewGrpcServerStream(stream grpc.ServerStream) io.StreamReader {
	return &serverStreamReader{}
}
