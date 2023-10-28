package main

import (
	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"github.com/boxcolli/go-transistor/cmd/mock"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	conn, err := grpc.Dial(mock.Addr, opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewTransistorServiceClient(conn)

	// client.Publish()
}
