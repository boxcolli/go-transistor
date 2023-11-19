package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"github.com/boxcolli/go-transistor/types"
	"github.com/peterbourgon/ff/v4"
	"google.golang.org/grpc"
)

func main() {
	// Parse flag, env var
	fs := flag.NewFlagSet("myprogram", flag.ContinueOnError)
	var (
		addr = fs.String("addr", ":443", "listen address")
		st = fs.String("st", "A0", "static 1st level topic for subscription")
	)
	ff.Parse(fs, os.Args[1:],
		ff.WithEnvVars(),
	)

	// Client
	var client pb.TransistorServiceClient
	{
		conn, err := grpc.Dial(*addr, dialOpts...)
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		client = pb.NewTransistorServiceClient(conn)
		fmt.Println("connected")
	}

	// Subscribe
	{
		var opts = []grpc.CallOption{}
		stream, err := client.Subscribe(context.Background(), opts...)
		if err != nil {
			panic(err)
		}

		// Send initial change
		cg := types.Change{ Op: types.OperationAdd, Topic: types.Topic{*st} }
		err = stream.Send(&pb.SubscribeRequest{
			Change: cg.Marshal(),
		})
		if err != nil {
			panic(err)
		}

		fmt.Println("now listening..")
		for {
			res, err := stream.Recv()
			if err != nil {
				log.Fatalf("Subscribe() received error: %v\n", err)
				break
			}

			msg := new(types.Message)
			msg.Unmarshal(res.GetMsg())
			log.Printf("Subscribe() receivd: %s\n", msg.String())
		}
	}
}
