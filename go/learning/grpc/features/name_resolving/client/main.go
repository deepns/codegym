package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	myresolver "github.com/deepns/codegym/go/learning/grpc/features/name_resolving/resolver"
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
		log.Fatalf("Failed to run UnaryEcho. err=%v", err)
	}

	log.Printf("echo: %v", response.Message)
}

func callRPCs(cc *grpc.ClientConn) {
	client := pb.NewEchoServiceClient(cc)
	UnaryEcho(client, cc.Target())
}

// dialWithDNSResolver connects to the server using the DNS resolver.
func dialWithDNSResolver(addr string) {
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server. err=%v", err)
	}
	defer conn.Close()
	callRPCs(conn)
}

// dialWithPassthroughResolver connects to the server using the passthrough resolver.
func dialWithPassthroughResolver(addr string) {
	passthroughConn, err := grpc.Dial(
		"passthrough:///"+addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server from passthrough client. err=%v", err)
	}
	defer passthroughConn.Close()
	callRPCs(passthroughConn)

}

// dialWithFooResolver connects to the server using the foo resolver.
func dialWithFooResolver(service string) {
	fooConn, err := grpc.Dial(service,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// This will register the FooResolverBuilder with the grpc resolver
		// only for the foo scheme. Overrides any existing resolver registered
		// with the global grpc resolver.
		grpc.WithResolvers(&myresolver.FooResolverBuilder{}))
	if err != nil {
		log.Fatalf("Failed to connect to server from foo client. err=%v", err)
	}
	defer fooConn.Close()
	callRPCs(fooConn)
}

func main() {
	addr := flag.String("addr", "localhost:50505", "address to connect to")
	flag.Parse()

	dialWithDNSResolver(*addr)

	dialWithPassthroughResolver(*addr)

	// Connecting to foo://resolver.foo.bar will fail unless we register
	// a resolver that resolves foo://resolver.foo.bar to localhost:50505.
	// Absence of a resolver for the foo scheme will result in a transport error.
	// Here is a sample error
	// err=rpc error: code = Unavailable desc = connection error: desc = "transport: Error while dialing dial tcp: lookup tcp////resolver.foo.bar: nodename nor servname provided, or not known"
	dialWithFooResolver("foo:///resolver.foo.bar")
}

// func init() {
// 	// This is another way to register the resolver.
// 	// For connection specific resolvers, use grpc.Dial with grpc.WithResolvers.
// 	resolver.Register(&myresolver.FooResolverBuilder{})
// }
