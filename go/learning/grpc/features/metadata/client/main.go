package main

import (
	"context"
	"flag"
	"log"
	"strings"
	"time"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func UnaryEcho(client pb.EchoServiceClient, message string) {
	request := pb.EchoRequest{
		Message: message,
	}

	// Adding some metadata to the context
	md := metadata.Pairs("timestamp", time.Now().Format(time.StampNano))
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	var header, trailer metadata.MD
	// pass header and trailer in the call options
	response, err := client.UnaryEcho(ctx, &request, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		log.Fatalf("Failed to run UnaryEcho. err=%v", err)
	}

	// Dump the header and trailer
	for k, v := range header {
		log.Printf("header: %v=%v", k, strings.Join(v, ","))
	}

	for k, v := range trailer {
		log.Printf("trailer: %v=%v", k, strings.Join(v, ","))
	}

	log.Printf("echo: %v", response.Message)
}

func main() {
	addr := flag.String("addr", "localhost:50505", "address to connect to")

	conn, err := grpc.Dial(*addr, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials())}...)
	if err != nil {
		log.Fatalf("Failed to connect to server. err=%v", err)
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)
	UnaryEcho(client, "Hello World")
}
