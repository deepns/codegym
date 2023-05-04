package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	"github.com/deepns/codegym/go/learning/grpc/features/sslcerts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type echoServer struct {
	pb.UnimplementedEchoServiceServer
}

func (s *echoServer) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("UnaryEcho: %v", req.Message)
	return &pb.EchoResponse{Message: req.Message}, nil
}

func main() {
	port := flag.Int("port", 50505, "port to listen")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	cert, err := credentials.NewServerTLSFromFile(
		sslcerts.Path("server_cert.pem"),
		sslcerts.Path("server_key.pem"))
	if err != nil {
		log.Fatalf("failed to load TLS cert: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(cert))
	pb.RegisterEchoServiceServer(s, &echoServer{})

	log.Printf("Listening on port %d", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
