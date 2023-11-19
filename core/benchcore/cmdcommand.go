package benchcore

import (
	"context"
	"crypto/tls"
	"flag"

	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"github.com/boxcolli/go-transistor/tools"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *benchCore) cmdCommand(ctx context.Context, args []string) (<-chan string, error) {
	fs := flag.NewFlagSet("command", flag.ContinueOnError)
	var (
		withTls = fs.Bool("tls", false, "grpc secure option")
		addr	= fs.String("addr", "", "target server address")
	)
	{
		err := fs.Parse(args)
		if err != nil { return nil, ErrInvalidArgument }
	}

	var opts []grpc.DialOption
	{
		if *withTls {
			opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
		} else {
			opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})))
		}
	}

	var client pb.TransistorServiceClient
	{
		var (
			conn *grpc.ClientConn
			err error
		)
		client, conn, err = tools.NewClient(*addr, opts)
		if err != nil { return nil, ErrUnavailable }
		defer conn.Close()
	}

	var stream pb.TransistorService_CommandClient
	{
		var err error
		stream, err = client.Command(ctx, &pb.CommandRequest{ Args: fs.Args() })
		if err != nil { return nil, ErrUnavailable }
	}

	var out = make(chan string, 10)
	go func() {
		defer close(out)

		for {
			res, err := stream.Recv()
			if err != nil { return }

			out <- res.GetLine()
		}
	} ()

	return out, nil
}
