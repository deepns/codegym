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

type myCredential struct{}

func (c myCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"username": "admin",
		"password": "secret",
	}, nil
}

func (c myCredential) RequireTransportSecurity() bool { return true }

func callUnaryEcho(c pb.EchoServiceClient, message string) {
	log.Printf("UnaryEcho: message=%v", message)
	request := pb.EchoRequest{
		Message: message,
	}

	response, err := c.UnaryEcho(context.Background(), &request)
	if err != nil {
		log.Fatalf("Failed to run UnaryEcho. err=%v", err)
	}

	log.Printf("UnaryEcho: echo_message=%v", response.Message)
}

func main() {
	addr := flag.String("addr", "localhost:50505", "address to connect to")
	flag.Parse()

	creds, err := credentials.NewClientTLSFromFile(sslcerts.Path("ca_cert.pem"), "abc.test.example.com")
	if err != nil {
		log.Fatalf("failed to load TLS cert: %v", err)
	}

	conn, err := grpc.Dial(*addr,
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(myCredential{}))

	if err != nil {
		log.Fatalf("Failed to connect to server. err=%v", err)
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)
	callUnaryEcho(client, "hello from static auth client")
}
