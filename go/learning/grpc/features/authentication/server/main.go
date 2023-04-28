package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"path"
	"strings"

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

func validateToken(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}

	token := authorization[0]

	// token comes in the format of "Bearer <token>"
	// so strip the prefix "Bearer "
	token = strings.TrimPrefix(token, "Bearer ")
	log.Printf("token:%v", token)

	// Just validating the dummy token
	return token == "my-oauth-token"
}

// UnaryInterceptor is a server interceptor function that
// intercepts and authenticates unary RPC with OAuth2 token.
func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("UnaryInterceptor: %v", info.FullMethod)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("unable to get metadata from the context")
	}

	for k, v := range md {
		log.Printf("metadata: %v=%v", k, strings.Join(v, ","))
	}

	if !validateToken(md["authorization"]) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
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

	serverOptions := []grpc.ServerOption{
		grpc.Creds(cert),                        // Enable TLS for all incoming connections. needed for authentication
		grpc.UnaryInterceptor(UnaryInterceptor), // interceptor for authentication
	}

	s := grpc.NewServer(serverOptions...)
	pb.RegisterEchoServiceServer(s, &echoServer{})

	log.Printf("Listening on port %d", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
