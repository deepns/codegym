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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type echoServer struct {
	pb.UnimplementedEchoServiceServer
}

func (s *echoServer) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("UnaryEcho: %v", req.Message)
	return &pb.EchoResponse{Message: req.Message}, nil
}

func authenticate(username string, password string) bool {
	if username != "admin" || password != "secret" {
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

	usernames, password := md.Get("username"), md.Get("password")
	if len(usernames) != 1 || len(password) != 1 {
		return nil, status.Errorf(codes.Unauthenticated, "authentication info missing")
	}

	if !authenticate(usernames[0], password[0]) {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed")
	}

	return handler(ctx, req)
}

func getCredentials() credentials.TransportCredentials {

	certPath := path.Join("../", "sslcerts")
	certFile := path.Join(certPath, "server_cert.pem")
	keyFile := path.Join(certPath, "server_key.pem")

	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	return creds
}

func main() {
	port := flag.Int("port", 50505, "port to listen")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	cert := getCredentials()

	s := grpc.NewServer(
		grpc.Creds(cert),
		grpc.UnaryInterceptor(unaryInterceptor))
	pb.RegisterEchoServiceServer(s, &echoServer{})

	log.Printf("Listening on port %d", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
