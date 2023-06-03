package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding/gzip" // installs the gzip compressor in its init() function
)

func callUnaryEcho(client pb.EchoServiceClient, msg string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use grpc.UseCompressor(gzip.Name) to enable gzip compression for this call
	resp, err := client.UnaryEcho(ctx, &pb.EchoRequest{Message: msg})
	if err != nil {
		log.Fatalf("UnaryEcho(%q) = %v", msg, err)
	}
	log.Printf("UnaryEcho(%q) = %q", msg, resp.Message)
}

func main() {
	addr := flag.String("addr", "localhost:50505", "address to connect to")
	flag.Parse()

	conn, err := grpc.Dial(*addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name))) // Use gzip compression for all calls
	if err != nil {
		log.Fatalf("Failed to connect to server. err=%v", err)
	}

	client := pb.NewEchoServiceClient(conn)
	callUnaryEcho(client, "hello from compression client")
}
