package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	"google.golang.org/grpc"
)

type echoServer struct {
	pb.UnimplementedEchoServiceServer
}

func (s *echoServer) SimpleEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	// Just a simple echo.
	// Send the request message back to the client
	return &pb.EchoResponse{Message: req.Message}, nil
}

func (s *echoServer) ServerSideStreamEcho(req *pb.EchoRequestWithCount, stream pb.EchoService_ServerSideStreamEchoServer) error {
	// Echo as many times as requested in the request

	var i int32
	for i = 0; i < req.Count; i++ {
		if err := stream.Send(&pb.EchoResponse{Message: req.Message}); err != nil {
			log.Printf("Failed to send. stream_id=%v, err=%v", i, err)
			return err
		}
	}
	return nil
}

func main() {
	port := flag.Int("port", 50505, "port to listen to")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", *port))
	if err != nil {
		log.Fatalf("Failed to listen at %v. err=%v", *port, err)
	}

	var options []grpc.ServerOption
	grpcServer := grpc.NewServer(options...)

	// register the server with grpc
	pb.RegisterEchoServiceServer(grpcServer, &echoServer{})

	// serve on the created listener
	grpcServer.Serve(lis)
}
