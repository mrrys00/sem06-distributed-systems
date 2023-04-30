package main

import (
	"grpcproject/grpcproject"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// based on https://itnext.io/build-grpc-server-with-golang-go-step-by-step-b3f5abcf9e0e
	println("gRPC server tutorial in Go")

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	grpcproject.RegisterGrpcProjectServer(s, &grpcproject.Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
