package main

import (
	"context"
	"flag"
	"io"
	"log"
	"math/rand"
	"strings"
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

func StreamMessages(client pb.EchoServiceClient, messages []string) {
	// timeout if response is not received within two seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	stream, err := client.ClientSideStreamEcho(ctx)
	if err != nil {
		log.Fatalf("client.ClientSideStreamEcho failed: %v", err)
	}

	for _, message := range messages {
		if err := stream.Send(&pb.EchoRequest{Message: message}); err != nil {
			log.Fatalf("stream.Send() failed: %v", err)
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("stream.CloseAndRecv() failed: %v", err)
	}

	log.Println("echo:", response.Message)
}

func randomStrings() []string {
	// Step 1: Pick a random number between 1..100
	rand.Seed(time.Now().UnixNano())
	numStrings := rand.Intn(10) + 1

	// Step 2: Create a slice of strings
	stringSlice := make([]string, numStrings)

	// Step 3: Fill the slice with random strings
	for i := 0; i < numStrings; i++ {
		// Generate a random string of length 5 using alphabets
		const letters = "abcdefghijklmnopqrstuvwxyz"
		b := make([]byte, 5)
		for j := range b {
			b[j] = letters[rand.Intn(len(letters))]
		}
		stringSlice[i] = string(b)
	}

	// Step 4: Return the slice
	return stringSlice
}

func ChitChat(client pb.EchoServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	stream, err := client.ChatEcho(ctx)
	if err != nil {
		log.Fatalf("client.ChatEcho failed: %v", err)
	}

	for _, message := range randomStrings() {
		log.Printf("send: %v", message)
		if err = stream.Send(&pb.EchoRequest{Message: message}); err != nil {
			log.Fatalf("stream.Send() failed: %v", err)
		}

		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream.Recv() failed: %v", err)
		}
		log.Printf("recv: %v", resp.Message)
	}
}

func main() {
	addr := flag.String("addr", "localhost:50505", "address of the server to connect to")
	message := flag.String("msg", "hello", "message to be sent to the echo server")
	stream := flag.String("stream", "", "comma separated list of messages to streamed to server")
	count := flag.Int("count", 1, "number of times to echo")
	chitchat := flag.Bool("chitchat", false, "exchange random messages through bidirectional streaming rpc")
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

	if *chitchat {
		ChitChat(client)
	} else if *count == 1 {
		if len(*stream) > 0 {
			StreamMessages(client, strings.Split(*stream, ","))
		} else {
			SimpleEcho(client, *message)
		}
	} else {
		EchoMultiple(client, *message, *count)
	}
}
