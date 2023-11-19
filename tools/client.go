package tools

import (
	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"

	"google.golang.org/grpc"
)

func NewClient(addr string, opts []grpc.DialOption) (
	client pb.TransistorServiceClient,
	conn *grpc.ClientConn,
	err error,
) {
	conn, err = grpc.Dial(addr, opts...)
	if err != nil {
		return
	}
	client = pb.NewTransistorServiceClient(conn)

	return
}
