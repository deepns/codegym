package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/deepns/codegym/go/learning/grpc/helloworld/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SayHello(client pb.HelloServiceClient, name string) {
	request := pb.HelloRequest{Name: name, Count: 2}
	// timeout the call after 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := client.SayHello(ctx, &request)
	if err != nil {
		log.Fatalf("Failed to run SayHello. err=%v", err)
	}
	log.Printf("Server sent: %v", response.HelloMsg)
}

func main() {
	serverAddr := flag.String("addr", "localhost:40404", "host:port of the server addr to connect to")
	flag.Parse()

	// Not using tls, so using insecure credentials
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// Create a connection to the server
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("Failed to dial to %v. err=%v", *serverAddr, err)
	}
	defer conn.Close()

	// create a new client on the service
	client := pb.NewHelloServiceClient(conn)

	SayHello(client, "foobar")
}
