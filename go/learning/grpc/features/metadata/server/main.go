package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type echoServer struct {
	pb.UnimplementedEchoServiceServer
}

/*
 * Some pending questions:
 * 1. What are the typical use cases for custom metadata?
 * 2. What information typically goes in the trailer?
 */

// Server side definition of the rpc method
func (s *echoServer) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	// Set the trailer to be set when the RPC is complete
	defer func() {
		trailer := metadata.Pairs("timestamp", time.Now().Format(time.StampNano))
		grpc.SetTrailer(ctx, trailer)
	}()

	// Read the metadata from the client
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "missing metadata")
	}

	for k, v := range md {
		log.Printf("metadata: %v=%v", k, strings.Join(v, ","))
	}

	header := metadata.New(map[string]string{"timestamp": time.Now().Format(time.StampNano)})
	grpc.SetHeader(ctx, header)

	// Note: SendHeader() can called at most once per RPC
	// Headers can also be sent by explicit call to SendHeader()
	// See SetHeader comments for other cases where headers are sent
	// (e.g when first message is returned or response is sent in case
	// of unary RPCs)
	// grpc.SendHeader(ctx, header)

	log.Printf("Request received: %v", req.Message)

	return &pb.EchoResponse{Message: req.Message}, nil
}

func main() {
	port := flag.Int("port", 50505, "port to listen to")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", *port))
	if err != nil {
		log.Fatalf("Failed to listen at %v. err=%v", *port, err)
	}

	server := grpc.NewServer()

	pb.RegisterEchoServiceServer(server, &echoServer{})
	server.Serve(lis)
}
