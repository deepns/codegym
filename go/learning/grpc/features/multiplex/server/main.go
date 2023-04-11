package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pbec "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	pbhw "github.com/deepns/codegym/go/learning/grpc/helloworld/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type echoServer struct {
	pbec.UnimplementedEchoServiceServer
}

type helloServer struct {
	pbhw.UnimplementedHelloServiceServer
}

func (s *echoServer) UnaryEcho(ctx context.Context, req *pbec.EchoRequest) (*pbec.EchoResponse, error) {
	log.Printf("UnaryEcho: %v", req.Message)
	if client, ok := peer.FromContext(ctx); ok {
		log.Printf("client:%v", client.Addr)
	}

	return &pbec.EchoResponse{Message: req.Message}, nil
}

func (s *helloServer) SayHello(ctx context.Context, req *pbhw.HelloRequest) (*pbhw.HelloResponse, error) {
	log.Printf("UnaryHello: %v", req.Name)
	if client, ok := peer.FromContext(ctx); ok {
		log.Printf("client:%v", client.Addr)
	}
	return &pbhw.HelloResponse{HelloMsg: fmt.Sprintf("Hello, %v", req.Name)}, nil
}

func main() {
	port := flag.Int("port", 50505, "port to connect to")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pbec.RegisterEchoServiceServer(s, &echoServer{})
	pbhw.RegisterHelloServiceServer(s, &helloServer{})

	log.Printf("Listening on port %d", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
