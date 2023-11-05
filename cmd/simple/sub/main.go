package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

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
		to = fs.Duration("to", 10 * time.Second, "grpc request timeout")
		st = fs.String("st", "A0", "static 1st level topic for subscription")
	)
	ff.Parse(fs, os.Args[1:],
		ff.WithEnvVars(),
	)


	// PubSub
	


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

	// Subscribe
	{
		var opts = []grpc.CallOption{}
		ctx, cancel := contextWithTimeout(*to)
		go func () {
			time.Sleep(*to)
			cancel()
		} ()
		stream, err := client.Subscribe(ctx, opts...)
		if err != nil {
			panic(err)
		}

		cg := types.Change{ Op: types.OperationAdd, Topic: types.Topic{*st} }
		err = stream.Send(&pb.SubscribeRequest{
			Change: cg.Marshal(),
		})
		if err != nil {
			panic(err)
		}

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

func contextWithTimeout(to time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), to)
}
