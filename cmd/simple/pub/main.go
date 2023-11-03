package main

import (
	"context"
	"flag"
	"fmt"
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
		to = fs.Duration("to", 100 * time.Second, "grpc request timeout")
		rate = fs.Duration("rate", 1 * time.Second, "message publish rate")
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

	// Publish
	{
		var opts = []grpc.CallOption{}
		ctx, cancel := contextWithTimeout(*to)
		defer cancel()

		stream, err := client.Publish(ctx, opts...)
		if err != nil {
			panic(err)
		}

		iteration := 0
		for {
			select {
			case <-ctx.Done():
				log.Fatalf("context finished: %v\n", ctx.Err())
				return
			default:
				msg := types.Message{
					Topic: types.Topic{"A0"},
					Mode: types.ModeAny,
					Method: types.MethodEmpty,
					Data: []byte(fmt.Sprintf("%v", iteration)),
					TP: time.Now().UTC(),
				}
	
				err := stream.Send(&pb.PublishRequest{ Msg: msg.Marshal() })
				if err != nil {
					log.Fatalf("stream failed: %v\n", err)
					break
				}
				log.Printf("sent message: %v\n", msg)
	
				iteration++
				time.Sleep(*rate)
			}
		}
	}
}

func contextWithTimeout(to time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), to)
}
