package main

import (
	"context"
	"flag"
	"log"
	"time"

	pbec "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	pbhw "github.com/deepns/codegym/go/learning/grpc/helloworld/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func callUnaryEcho(client pbec.EchoServiceClient, message string) {
	request := pbec.EchoRequest{
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

	log.Printf("response from echo service: %v", response.Message)
}

func callSayHello(client pbhw.HelloServiceClient, name string) {
	request := pbhw.HelloRequest{
		Name:  name,
		Count: 2,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	response, err := client.SayHello(ctx, &request)
	if err != nil {
		log.Fatalf("Failed to run SayHello. err=%v", err)
	}

	log.Printf("response from hello service: %v", response.HelloMsg)
}

func main() {
	addr := flag.String("addr", "localhost:50505", "address to connect to")
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server. err=%v", err)
	}
	defer conn.Close()

	echoClient := pbec.NewEchoServiceClient(conn)
	callUnaryEcho(echoClient, "Boo!!!")

	helloClient := pbhw.NewHelloServiceClient(conn)
	callSayHello(helloClient, "Steve")
}
