package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/peer"
)

type echoServer struct {
	pb.UnimplementedEchoServiceServer
}

var enforcementPolicy = keepalive.EnforcementPolicy{
	MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
	PermitWithoutStream: true,            // Allow pings even when there are no active streams
}

var serverParams = keepalive.ServerParameters{
	MaxConnectionIdle:     10 * time.Second, // If a client is idle for 10 seconds, send a GOAWAY
	MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
	MaxConnectionAgeGrace: 10 * time.Second, // Allow 10 seconds for pending RPCs to complete before forcibly closing connections
	Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
}

func (s *echoServer) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	client, ok := peer.FromContext(ctx)
	if !ok {
		log.Println("Unable to get peer info from the context")
	}
	log.Printf("client:%v, msg:%v", client.Addr, req.Message)

	// Just a simple echo.
	// Send the request message back to the client
	return &pb.EchoResponse{Message: req.Message}, nil
}

func main() {
	port := flag.Int("port", 50055, "The server port")
	flag.Parse()

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", *port))
	if err != nil {
		log.Fatalf("Failed to listen on port %v. err=%v", *port, err)
	}

	// Add keepalive options to the server
	var options = []grpc.ServerOption{
		grpc.KeepaliveEnforcementPolicy(enforcementPolicy),
		grpc.KeepaliveParams(serverParams),
	}

	// Create a server instance
	grpcServer := grpc.NewServer(options...)

	// Attach the echo service to the server
	pb.RegisterEchoServiceServer(grpcServer, &echoServer{})

	// Serve and block
	grpcServer.Serve(lis)
}
