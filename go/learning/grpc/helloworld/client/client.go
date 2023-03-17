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

func SayHello(client pb.HelloServiceClient, name string, repeat uint32) {
	request := pb.HelloRequest{Name: name, Count: repeat}

	// timeout the call after 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// grpc generated client SayHello takes a context and request as inputs
	response, err := client.SayHello(ctx, &request)
	if err != nil {
		log.Fatalf("Failed to run SayHello. err=%v", err)
	}
	log.Printf("Server sent: %v", response.HelloMsg)
}

func main() {
	serverAddr := flag.String("addr", "localhost:40404", "host:port of the server addr to connect to")
	name := flag.String("name", "foo", "Name to say hello")
	repeat := flag.Uint("repeat", 3, "Number of times to repeat")

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

	SayHello(client, *name, uint32(*repeat))
}
