package grpcwriter

import (
	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

type grpcWriter struct {
}

// Write implements io.StreamWriter.
func (*grpcWriter) Write(*types.Message) error {
	panic("unimplemented")
}

func NewGrpcWriter(stream pb.TransistorService_SubscribeServer) io.StreamWriter {
	return &grpcWriter{}
}
