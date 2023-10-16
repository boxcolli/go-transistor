package grpcreader

import (
	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
	"google.golang.org/grpc"
)

type clientStreamReader struct {
	c pb.TransistorService_SubscribeClient
}

// Read implements StreamReader.
func (r *clientStreamReader) Read() (*types.Message, error) {
	// Pull new message
	res, err := r.c.Recv()
	if err != nil {
		return nil, err
	}
	
	// Unmarshal the message
	m := new(types.Message)
	m.Unmarshal(res.GetMsg())
	return m, nil
}

func NewGrpcClientStream(c pb.TransistorService_SubscribeClient) io.StreamReader {
	return &clientStreamReader{ c: c }
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
