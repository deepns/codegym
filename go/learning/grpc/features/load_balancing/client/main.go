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

// UnaryEcho sends a unary RPC to the server.
func UnaryEcho(client pb.EchoServiceClient, message string, callCount int) {
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

	log.Printf("echo(%v): %v", callCount, response.Message)
}

func callRPCs(cc *grpc.ClientConn, count int) {
	client := pb.NewEchoServiceClient(cc)
	for i := 0; i < count; i++ {
		UnaryEcho(client, cc.Target(), i)
	}
}

// dialWithRoundRobin connects to the service using the round robin load balancing policy.
func dialWithRoundRobin(service string, rpcCount int) {
	// gRPC supports two types of load balancing policies on the client side:
	// 1. PickFirst: The client will connect to the first server that it resolves.
	// 2. RoundRobin: The client will connect to the servers in a round robin fashion.
	// The default load balancing policy is PickFirst.
	// To use the RoundRobin load balancing policy, we need to set the default service config.
	// The default service config is a JSON string that contains the load balancing policy.
	// With round robin, client connects to all addresses it sees and sends RPC to each
	// backend address one at a time. If the connection is not ready for any reason, client
	// may send RPCs to the same backend address multiple times.
	fooConn, err := grpc.Dial(service,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		log.Fatalf("Failed to connect to server from foo client. err=%v", err)
	}
	defer fooConn.Close()
	callRPCs(fooConn, rpcCount)
}

func main() {
	addressStr := flag.String("addr", "localhost:50505", "comma separated addresses to connect to")
	rpcCount := flag.Int("rpc_count", 10, "number of RPCs to make")
	flag.Parse()

	myresolver.ServerAddresses = strings.Split(*addressStr, ",")
	dialWithRoundRobin("foo:///resolver.foo.bar", *rpcCount)
}

func init() {
	// register the foo resolver with the grpc resolver
	resolver.Register(&myresolver.FooResolverBuilder{})
}
