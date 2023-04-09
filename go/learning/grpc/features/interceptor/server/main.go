package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type echoServer struct {
	pb.UnimplementedEchoServiceServer
}

func (s *echoServer) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	peer, ok := peer.FromContext(ctx)
	if !ok {
		log.Println("Unable to get peer info from the context")
	}
	log.Printf("peer:%v", peer.Addr)

	// Just a simple echo.
	// Send the request message back to the client
	return &pb.EchoResponse{Message: req.Message}, nil
}

// unaryInterceptor intercepts the incoming request and logs the request
func unaryInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	// Do some pre-processing
	// FullMethod is the full RPC method string, i.e., /package.service/method.
	// In this case, it will be /echo.EchoService/UnaryEcho
	log.Printf("Unary interceptor: %v", info.FullMethod)
	log.Printf("Request: %v", req)

	// Read the metadata from the client
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("Unable to get metadata from the context")
	} else {
		for k, v := range md {
			log.Printf("Metadata: %v=%v", k, strings.Join(v, ","))
		}
	}

	// Sample output of the above log statements:
	// 2023/04/09 08:07:30 Unary interceptor: /EchoService/UnaryEcho
	// 2023/04/09 08:07:30 Request: message:"hello"
	// 2023/04/09 08:07:30 Metadata: :authority=localhost:50505
	// 2023/04/09 08:07:30 Metadata: content-type=application/grpc
	// 2023/04/09 08:07:30 Metadata: user-agent=grpc-go/1.52.0

	// Call the handler and return the response
	return handler(ctx, req)
}

func main() {
	port := flag.Int("port", 50505, "port to listen to")
	flag.Parse()

	// when just 'tcp' is used, it will listen on all interfaces
	// use 'tcp4' or 'tcp6' to listen on specific interfaces
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		log.Fatalf("Failed to listen at %v. err=%v", *port, err)
	}

	server := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor))
	pb.RegisterEchoServiceServer(server, &echoServer{})
	server.Serve(lis)
}
