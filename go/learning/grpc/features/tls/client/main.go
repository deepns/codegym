package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	"github.com/deepns/codegym/go/learning/grpc/features/sslcerts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func callUnaryEcho(client pb.EchoServiceClient, msg string) {
	request := pb.EchoRequest{
		Message: msg,
	}

	ctx := context.Background()
	response, err := client.UnaryEcho(ctx, &request)
	if err != nil {
		log.Fatalf("Failed to call UnaryEcho. err=%v", err)
	}

	log.Printf("echo: %v", response.Message)
}

func main() {
	addr := flag.String("addr", "localhost:50505", "address to connect to")
	flag.Parse()

	creds, err := credentials.NewClientTLSFromFile(sslcerts.Path("ca_cert.pem"), "abc.test.example.com")
	if err != nil {
		log.Fatalf("failed to load TLS cert: %v", err)
	}

	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}
	conn, err := grpc.Dial(*addr, dialOptions...)
	if err != nil {
		log.Fatalf("Failed to connect to server. err=%v", err)
	}

	client := pb.NewEchoServiceClient(conn)
	callUnaryEcho(client, "hello from tls client")
}
