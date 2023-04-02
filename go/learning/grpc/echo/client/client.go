package main

import (
	"bufio"
	"context"
	"flag"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UnaryEcho(client pb.EchoServiceClient, message string) {
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

	log.Printf("echo: %v", response.Message)
}

func ServerSideStream(client pb.EchoServiceClient, message string, count int) {
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

func ClientSideStream(client pb.EchoServiceClient, messages []string) {
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

// randomStringSlice generates a slice of random words selected from a dictionary.
// The number of random words generated is determined by a random integer between 1 and 100.
// The function returns a slice of the randomly generated words.
func randomStringSlice(count int) []string {
	// Step 1: Pick a random number between 1..100
	rand.Seed(time.Now().UnixNano())
	numStrings := rand.Intn(count) + 1

	// Step 2: Create a slice of strings
	stringSlice := make([]string, numStrings)

	// Step 3: Fill the slice with random strings
	wordList, err := os.Open("/usr/share/dict/words") // change path to your system's word list file
	if err != nil {
		panic(err)
	}
	defer wordList.Close()
	scanner := bufio.NewScanner(wordList)
	words := []string{}
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	for i := 0; i < numStrings; i++ {
		// Generate a random word from the dictionary
		stringSlice[i] = words[rand.Intn(len(words))]
	}

	// Step 4: Return the slice
	return stringSlice
}

func BidirectionalStream(client pb.EchoServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	stream, err := client.BidrectionalStreamEcho(ctx)
	if err != nil {
		log.Fatalf("client.ChatEcho failed: %v", err)
	}

	for _, message := range randomStringSlice(10) {
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
	var addr, rpc string

	flag.StringVar(&addr, "addr", "localhost:50505", "address of the server to connect to")
	flag.StringVar(&rpc, "rpc", "", "Specify the RPC value (valid options: unary, clientstream, serverstream, bidirectional)")
	flag.Parse()

	// Not using TLS. So using insecure credentials
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// Create a connection to the server
	conn, err := grpc.Dial(addr, options...)
	if err != nil {
		log.Fatalf("Failed to connect to the server %v. err=%v", addr, err)
	}
	defer conn.Close()

	// Create a new client to the chosen service
	client := pb.NewEchoServiceClient(conn)

	if rpc == "unary" {
		UnaryEcho(client, "test-unary-echo")
	} else if rpc == "clientstream" {
		ClientSideStream(client, []string{
			"test-clientstream-msg-1",
			"test-clientstream-msg-2",
			"test-clientstream-msg-3",
		})
	} else if rpc == "serverstream" {
		ServerSideStream(client, "test-serverstream-msg", 5)
	} else if rpc == "bidirectional" {
		BidirectionalStream(client)
	} else {
		log.Fatalf("Invalid RPC: %v", rpc)
	}
}
