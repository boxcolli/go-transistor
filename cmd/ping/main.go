package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"github.com/peterbourgon/ff/v4"
	"google.golang.org/grpc"
)

func main() {
	// Parse flag, env var
	fs := flag.NewFlagSet("myprogram", flag.ContinueOnError)
	var (
		addr = fs.String("addr", ":443", "listen address")
		to = fs.Duration("to", 10 * time.Second, "grpc request timeout")
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
	}

	// Ping
	{
		var opts = []grpc.CallOption{}
		ctx, cancel := contextWithTimeout(*to)
		go func () {
			time.Sleep(*to)
			cancel()
		} ()
		stream, err := client.Command(ctx, &pb.CommandRequest{ Args: []string{"ping"} }, opts...)
		if err != nil {
			panic(err)
		}

		for {
			res, err := stream.Recv()
			if err != nil {
				log.Fatalf("Command(ping) received error: %v\n", err)
				break
			}

			log.Printf("Command(ping) received: %s\n", res.GetLine())
		}
	}
}

func contextWithTimeout(to time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), to)
}
