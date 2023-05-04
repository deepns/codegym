package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	"github.com/deepns/codegym/go/learning/grpc/features/sslcerts"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
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

func getToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "my-oauth-token",
	}
}

func main() {
	addr := flag.String("addr", "localhost:50505", "address to connect to")
	flag.Parse()

	creds, err := credentials.NewClientTLSFromFile(sslcerts.Path("ca_cert.pem"), "abc.test.example.com")
	if err != nil {
		log.Fatalf("failed to load TLS cert: %v", err)
	}

	// Use this token for all RPC calls
	tokenSource := oauth.TokenSource{TokenSource: oauth2.StaticTokenSource(getToken())}

	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		// apply these credentials for all RPC calls
		grpc.WithPerRPCCredentials(tokenSource),
	}

	conn, err := grpc.Dial(*addr, dialOptions...)
	if err != nil {
		log.Fatalf("Failed to connect to server. err=%v", err)
	}

	client := pb.NewEchoServiceClient(conn)
	callUnaryEcho(client, "hello from tls client")
}
