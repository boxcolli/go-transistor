package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"github.com/boxcolli/go-transistor/cmd/mock"
	"github.com/boxcolli/go-transistor/server/mockserver"
	"google.golang.org/grpc"
)

const (
	chsiz = 10
)

func main() {
	port := os.Getenv("PORT")
        if port == "" {
                port = mock.Port
                log.Printf("Defaulting to port %s", port)
        }
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTransistorServiceServer(s, mockserver.NewMockServer(chsiz))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
