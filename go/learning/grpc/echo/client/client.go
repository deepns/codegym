package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SimpleEcho(client pb.EchoServiceClient, message string) {
	request := pb.EchoRequest{
		Message: message,
	}

	// timeout if response is not received within two seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// connect to the service
	response, err := client.SimpleEcho(ctx, &request)
	if err != nil {
		log.Fatalf("Failed to run SimpleEcho. err=%v", err)
	}

	log.Printf("echo: %v", response.Message)
}

func EchoMultiple(client pb.EchoServiceClient, message string, count int) {
	log.Printf("Echoing %v, %v times", message, count)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	stream, err := client.ServerSideStreamEcho(ctx, &pb.EchoRequestWithCount{
		Message: message,
		Count:   int32(count),
	})

	if err != nil {
		log.Fatalf("client.ServerSideStreamEcho failed. err=%v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to recv from stream. err=%v", err)
		}
		log.Printf("echo: %v", resp.Message)
	}
}

func main() {
	addr := flag.String("addr", "localhost:50505", "address of the server to connect to")
	message := flag.String("msg", "foobar", "message to be sent to the echo server")
	count := flag.Int("count", 1, "number of times to echo")
	flag.Parse()

	if *count < 1 {
		log.Fatalf("Too low count to repeat. Must be between 1...10")
	}

	if *count > 10 {
		log.Fatalf("Too high to repeat Must be between 1...10")
	}

	// Not using TLS. So using insecure credentials
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// Create a connection to the server
	conn, err := grpc.Dial(*addr, options...)
	if err != nil {
		log.Fatalf("Failed to connect to the server %v. err=%v", *addr, err)
	}
	defer conn.Close()

	// Create a new client to the chosen service
	client := pb.NewEchoServiceClient(conn)

	if *count == 1 {
		SimpleEcho(client, *message)
	} else {
		EchoMultiple(client, *message, *count)
	}
}
