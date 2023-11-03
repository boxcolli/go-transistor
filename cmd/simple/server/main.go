package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"github.com/boxcolli/go-transistor/base/basicbase"
	"github.com/boxcolli/go-transistor/collector/basiccollector"
	"github.com/boxcolli/go-transistor/core"
	"github.com/boxcolli/go-transistor/core/simplecore"
	"github.com/boxcolli/go-transistor/index/routeindex"
	server "github.com/boxcolli/go-transistor/server/grpcserver"
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
		eqs = fs.Int("eqs", 100, "emitter queue size")
	)
	ff.Parse(fs, os.Args[1:],
		ff.WithEnvVars(),
	)

	// Transistor
	var core core.Core
	{
		base := basicbase.NewBasicBase(routeindex.NewRouteIndex, *bcqs)
		collector := basiccollector.NewBasicCollector(base, *cmqs)
		core = simplecore.NewSimpleCore(simplecore.Component{
			Base: base,
			Collector: collector,
			}, simplecore.Option{})
		core.Start()
		log.Println("core started.")
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
		pb.RegisterTransistorServiceServer(grpcServer, server.NewTransistorServer(core, *eqs))
		logger.Printf("server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			logger.Fatalf("failed to serve: %v", err)
		}
	}
}
