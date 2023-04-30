package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"path"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type echoServer struct {
	pb.UnimplementedEchoServiceServer
}

func (s *echoServer) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("UnaryEcho: %v", req.Message)
	return &pb.EchoResponse{Message: req.Message}, nil
}

func authenticate(username string, password string) bool {
	if username != "admin" || password != "password" {
		return false
	}

	return true
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("unable to get metadata from the context")
	}

	for k, v := range md {
		log.Printf("metadata: %v=%v", k, v)
	}

	// What if username is missing?
	//
	username := md.Get("username")[0]
	password := md.Get("password")[0]

	if !authenticate(username, password) {
		return nil, fmt.Errorf("no auth info")
	}

	if !authenticate(username, password) {
		return nil, fmt.Errorf("authentication failed")
	}

	return handler(ctx, req)
}

func main() {
	port := flag.Int("port", 50505, "port to listen")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	certPath := path.Join("..", "sslcerts")
	cert, err := credentials.NewServerTLSFromFile(
		path.Join(certPath, "server_cert.pem"),
		path.Join(certPath, "server_key.pem"))
	if err != nil {
		log.Fatalf("failed to load TLS cert: %v", err)
	}

	s := grpc.NewServer(
		grpc.Creds(cert),
		grpc.UnaryInterceptor(unaryInterceptor))
	pb.RegisterEchoServiceServer(s, &echoServer{})

	log.Printf("Listening on port %d", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
