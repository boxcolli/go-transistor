package server

import (
	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
)

type Server interface {
	pb.TransistorServiceServer
}
