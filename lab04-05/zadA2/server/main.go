package main

import (
	context "context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"grpcproject/grpcproject"

	"google.golang.org/grpc"
)

// We define a server struct that implements the server interface. 🥳🥳🥳
type server struct {
	grpcproject.UnimplementedGrpcProjectServer
}

// We implement the SayHello method of the server interface. 🥳🥳🥳
func (s *server) SayHello(ctx context.Context, in *grpcproject.HelloRequest) (*grpcproject.HelloReply, error) {
	fmt.Printf("Recived message: %s\n", in.GetName())
	return &grpcproject.HelloReply{Message: "Hello, " + in.GetName()}, nil
}

func (s *server) FetchResponse(in *grpcproject.Request, srv grpcproject.GrpcProject_FetchResponseServer) error {

	log.Printf("fetch response for id : %d", in.Id)

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(count int64) {		// 
			defer wg.Done()
			time.Sleep(time.Duration(count) * time.Second)
			resp := grpcproject.Response{Result: fmt.Sprintf("Request #%d For Id:%d", count, in.Id)}
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
			log.Printf("finishing request number : %d", count)
		}(int64(i))
	}

	wg.Wait()
	return nil
}

func main() {
	// based on https://itnext.io/build-grpc-server-with-golang-go-step-by-step-b3f5abcf9e0e
	println("gRPC server tutorial in Go")

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	grpcproject.RegisterGrpcProjectServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
