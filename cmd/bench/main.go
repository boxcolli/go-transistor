package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"github.com/boxcolli/go-transistor/base/basicbase"
	"github.com/boxcolli/go-transistor/collector/basiccollector"
	"github.com/boxcolli/go-transistor/transistor"
	"github.com/boxcolli/go-transistor/transistor/benchtransistor"
	"github.com/boxcolli/go-transistor/emitter/basicemitter"
	"github.com/boxcolli/go-transistor/index/basicindex"
	"github.com/boxcolli/go-transistor/server/grpcserver"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/peterbourgon/ff/v4"
	"google.golang.org/grpc"
)

func main() {
	// Parse flag, env var
	fs := flag.NewFlagSet("myprogram", flag.ContinueOnError)
	var (
		port = fs.String("port", "443", "listen port")
		cmqs = fs.Int("cmqs", 10, "collector message queue size")
		bcqs = fs.Int("bcqs", 100, "base change queue size")
		emqs = fs.Int("emqs", 10, "emitter message queue size")
	)
	ff.Parse(fs, os.Args[1:],
		ff.WithEnvVars(),
	)

	// Transistor
	var tr transistor.Transistor
	{
		tr = benchtransistor.NewBenchTransistor(
			transistor.Component{
				Base: basicbase.NewBasicBase(basicindex.NewBasicIndex, *bcqs),
				Collector: basiccollector.NewBasicCollector(*cmqs),
				Emitter: basicemitter.NewBasicEmitter(*emqs),
			},
			benchtransistor.Option{},
		)
		fmt.Println("tr started.")
	}

	// Server
	{
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
		if err != nil {
			logger.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				logging.UnaryServerInterceptor(InterceptorLogger(logger), logOpts...),
			),
			grpc.ChainStreamInterceptor(
				logging.StreamServerInterceptor(InterceptorLogger(logger), logOpts...),
			),
		)
		pb.RegisterTransistorServiceServer(grpcServer, grpcserver.NewGrpcServer(tr))
		logger.Printf("server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			logger.Fatalf("failed to serve: %v", err)
		}
	}
}
