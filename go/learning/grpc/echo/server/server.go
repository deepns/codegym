package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type echoServer struct {
	pb.UnimplementedEchoServiceServer
}

func (s *echoServer) SimpleEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	peer, ok := peer.FromContext(ctx)
	if !ok {
		log.Println("Unable to get peer info from the context")
	}
	log.Printf("peer:%v", peer.Addr)

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

func (s *echoServer) ClientSideStreamEcho(stream pb.EchoService_ClientSideStreamEchoServer) error {
	reqCount := 0
	var reqs []string
	for {
		req, err := stream.Recv()
		reqCount += 1
		if err == io.EOF || reqCount >= 10 {
			log.Println("No more incoming messages")
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive from stream: %v", err)
		}
		log.Printf("Received: %v", req.Message)
		reqs = append(reqs, req.Message)
	}

	return stream.SendAndClose(&pb.EchoResponse{Message: strings.Join(reqs, "::")})
}

func (s *echoServer) ChatEcho(stream pb.EchoService_ChatEchoServer) error {
	// TODO
	// Is there a way to identify client connection details from the stream?
	// Like IP or some unique identifier for a connection? - Yes. can do that
	// using peer package.
	peer, ok := peer.FromContext(stream.Context())
	if !ok {
		log.Println("Unable to get peer info from context")
	}
	log.Printf("peer: %v", peer.Addr)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			// If the client connection is closed, stream.Recv() may fail with
			// error "Canceled". That's ok.
			return err
		}

		log.Printf("received: %v", req.Message)
		if err = stream.Send(&pb.EchoResponse{Message: req.Message}); err != nil {
			log.Printf("stream.Send() failed: message: %v, err:%v", req.Message, err)
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
