package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var dialOpts = []grpc.DialOption{
	grpc.WithTransportCredentials(insecure.NewCredentials()),
}