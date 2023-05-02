package main

import (
	"context"
	"flag"
	"log"
	"path"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func callUnaryEcho(c pb.EchoServiceClient, message string, username string, password string) {
	// well, not a good practice to log password. but this is just an example.
	log.Printf("UnaryEcho: message=%v, username=%v, password=%v", message, username, password)
	request := pb.EchoRequest{
		Message: message,
	}

	ctx := metadata.NewOutgoingContext(context.Background(),
		metadata.Pairs("username", username, "password", password))

	response, err := c.UnaryEcho(ctx, &request)
	if err != nil {
		log.Fatalf("Failed to run UnaryEcho. err=%v", err)
	}

	log.Printf("UnaryEcho: echo_message=%v", response.Message)
}

func main() {
	addr := flag.String("addr", "localhost:50505", "address to connect to")
	flag.Parse()

	certPath := "../sslcerts"
	creds, err := credentials.NewClientTLSFromFile(path.Join(certPath, "ca_cert.pem"), "abc.test.example.com")
	if err != nil {
		log.Fatalf("failed to load TLS cert: %v", err)
	}

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect to server. err=%v", err)
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)
	callUnaryEcho(client, "hello from static auth client", "admin", "secret")
	callUnaryEcho(client, "hello from static auth client", "admin", "invalid-secret")
}
