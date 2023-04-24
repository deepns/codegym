package main

import (
	"context"
	"flag"
	"log"
	"strings"
	"time"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	myresolver "github.com/deepns/codegym/go/learning/grpc/features/name_resolving/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
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
		log.Fatalf("Failed to run UnaryEcho. err=%v", err)
	}

	log.Printf("echo: %v", response.Message)
}

func callRPCs(cc *grpc.ClientConn) {
	client := pb.NewEchoServiceClient(cc)
	UnaryEcho(client, cc.Target())
}

// dialWithRoundRobin connects to the service using the round robin load balancing policy.
func dialWithRoundRobin(service string) {
	fooConn, err := grpc.Dial(service,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		log.Fatalf("Failed to connect to server from foo client. err=%v", err)
	}
	defer fooConn.Close()
	callRPCs(fooConn)
}

func main() {
	addressStr := flag.String("addr", "localhost:50505", "comma separated addresses to connect to")
	flag.Parse()

	myresolver.ServerAddresses = strings.Split(*addressStr, ",")
	for range myresolver.ServerAddresses {
		dialWithRoundRobin("foo:///resolver.foo.bar")
	}

	// TODO
	// [ ] Verify that the client is using the round robin load balancing policy.
	// [ ] Check if we can get the resolved address from the client connection.
	// [ ] Add notes for the load balancing feature
	// [ ] Check what other load balancing policies are available.
}

func init() {
	// register the foo resolver with the grpc resolver
	resolver.Register(&myresolver.FooResolverBuilder{})
}
