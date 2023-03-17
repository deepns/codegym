package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/deepns/codegym/go/learning/grpc/helloworld/helloworld"
	"google.golang.org/grpc"
)

type helloServer struct {
	// must be embedded to have forward compatible implementation
	pb.UnimplementedHelloServiceServer
}

func (s *helloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{HelloMsg: fmt.Sprintf("Hello, %v", req.Name)}, nil
}

func main() {
	port := flag.Int("port", 40404, "port to listen")
	flag.Parse()

	lis, err := net.Listen("tcp", /* type of network socket */
		fmt.Sprintf("localhost:%v", *port))
	if err != nil {
		log.Fatalf("Failed to listen on %v. err=%v", *port, err)
	}

	// not using tls. so just an empty option would do.
	var options []grpc.ServerOption
	grpcServer := grpc.NewServer(options...)

	// register the server with the grpc service
	pb.RegisterHelloServiceServer(grpcServer, &helloServer{})

	// serve on the created listener
	grpcServer.Serve(lis)
}
