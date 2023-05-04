package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

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
		log.Fatalf("Failed to run SimpleEcho. err=%v", err)
	}

	log.Printf("UnaryEcho: echo_message=%v", response.Message)
}

func UnaryClientIntercept(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	// Pre-processing
	log.Printf("UnaryClientIntercept: method=%v", method)
	log.Printf("UnaryClientIntercept: req=%v", req)
	log.Printf("UnaryClientIntercept: ctx=%v", ctx)
	for opt := range opts {
		log.Printf("UnaryClientIntercept: opt=%v", opt)
	}

	// Can modify the request or call options too.
	// e.g. use case.
	// Add a header to the request.
	// Add authentication token to the request if it is not present.
	opts = append(opts, grpc.Header(&metadata.MD{"myheader": []string{"myvalue"}}))

	// Call the invoker
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)

	// Post-processing
	timeElapsed := time.Since(start)
	log.Printf("UnaryClientIntercept: reply=%v, err=%v, timeElapsed=%v", reply, err, timeElapsed)

	return err
}

func main() {
	addr := flag.String("addr", "localhost:50505", "address to connect to")
	msg := flag.String("msg", "hello", "message to echo")
	flag.Parse()

	conn, err := grpc.Dial(*addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(UnaryClientIntercept))

	if err != nil {
		log.Fatalf("Failed to connect to server. err=%v", err)
	}

	client := pb.NewEchoServiceClient(conn)
	callUnaryEcho(client, *msg)
}
