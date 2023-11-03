package grpcreader

import (
	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

type clientStreamReader struct {
	stream pb.TransistorService_SubscribeClient
}

// Read implements StreamReader.
func (r *clientStreamReader) Read() (*types.Message, error) {
	// Pull new message
	res, err := r.stream.Recv()
	if err != nil {
		return nil, err
	}
	
	// Unmarshal the message
	m := new(types.Message)
	m.Unmarshal(res.GetMsg())
	return m, nil
}

func NewGrpcClientStream(stream pb.TransistorService_SubscribeClient) io.StreamReader {
	return &clientStreamReader{ stream: stream }
}

type serverStreamReader struct {
	stream pb.TransistorService_PublishServer
}

// Read implements StreamReader.
func (r *serverStreamReader) Read() (*types.Message, error) {
	// Pull new message
	res,err := r.stream.Recv()
	if err != nil {
		return nil, err
	}

	// Unmarshal the message
	m := new(types.Message)
	m.Unmarshal(res.GetMsg())
	return m, nil
}

func NewGrpcServerStream(stream pb.TransistorService_PublishServer) io.StreamReader {
	return &serverStreamReader{ stream: stream }
}
