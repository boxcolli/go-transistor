package grpcwriter

import (
	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

type grpcWriter struct {
	stream pb.TransistorService_SubscribeServer 
}

// Write implements io.StreamWriter.
func (w *grpcWriter) Write(m *types.Message) error {
	return w.stream.Send(&pb.SubscribeResponse{
		Msg: m.Marshal(),
	})
}

func NewGrpcWriter(stream pb.TransistorService_SubscribeServer) io.StreamWriter {
	return &grpcWriter{
		stream: stream,
	}
}
