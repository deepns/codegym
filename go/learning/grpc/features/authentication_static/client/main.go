package main

import (
	"context"
	"flag"
	"log"
	"path"
	"time"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func unaryClientInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Printf("before call: %s, request: %+v", method, req)

	opts = append(opts, grpc.Header(&metadata.MD{"username": []string{"admin"}, "password": []string{"password"}}))

	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("after call: %s, response: %+v", method, reply)
	return err
}

func callUnaryEcho(client pb.EchoServiceClient, message string) {
	log.Printf("UnaryEcho: message=%v", message)
	request := pb.EchoRequest{
		Message: message,
	}

	// timeout if response is not received within two seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// connect to the service
	response, err := client.UnaryEcho(ctx, &request)
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

	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		grpc.WithUnaryInterceptor(unaryClientInterceptor)}

	conn, err := grpc.Dial(*addr, dialOptions...)
	if err != nil {
		log.Fatalf("Failed to connect to server. err=%v", err)
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)
	callUnaryEcho(client, "hello from static auth client")
}
