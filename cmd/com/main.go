package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"google.golang.org/grpc"
)

func main() {
	// Parse flag, env var
	fs := flag.NewFlagSet("main", flag.ContinueOnError)
	var (
		addr = fs.String("addr", ":443", "listen address")
	)
	{
		err := fs.Parse(os.Args[1:])
		if err != nil { panic(err) }
	}
	fmt.Println("request args:", fs.Args())

	// Client
	var client pb.TransistorServiceClient
	{
		conn, err := grpc.Dial(*addr, dialOpts...)
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		client = pb.NewTransistorServiceClient(conn)
	}

	var stream pb.TransistorService_CommandClient
	{
		var err error
		stream ,err = client.Command(context.Background(), &pb.CommandRequest{
			Args: fs.Args(),
		})
		if err != nil { panic(err) }
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			fmt.Printf("received err: %v\n", err)
			break
		}

		fmt.Println(res.GetLine())
	}
}
