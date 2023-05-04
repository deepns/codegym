package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	pbec "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	pbhw "github.com/deepns/codegym/go/learning/grpc/helloworld/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
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

	// Limit the number of times we say hello
	if req.Count > 10 {
		req.Count = 10
	}

	helloMsg := strings.Repeat(fmt.Sprintf("Hello, %v\n", req.Name), int(req.Count))
	return &pbhw.HelloResponse{HelloMsg: helloMsg}, nil
}

func main() {
	// This is pretty much the same as the server I used in multiplex example
	// except that I have added reflection to the server
	port := flag.Int("port", 50505, "port to connect to")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pbec.RegisterEchoServiceServer(s, &echoServer{})
	pbhw.RegisterHelloServiceServer(s, &helloServer{})

	reflection.Register(s)

	log.Printf("Listening on port %d", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
