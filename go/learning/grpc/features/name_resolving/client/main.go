package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
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

var (
	serviceName    = "resolver.foo.bar"
	serviceBackend = "localhost:50505"
)

func callRPCs(cc *grpc.ClientConn) {
	log.Printf("callRPCs: cc.Target()=%v", cc.Target())
	client := pb.NewEchoServiceClient(cc)
	UnaryEcho(client, "Call from name_resolving:client")
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
	fooConn, err := grpc.Dial(
		fmt.Sprintf("foo:///%s", service),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// This will register the fooResolverBuilder with the grpc resolver
		// only for the foo scheme. Overrides any existing resolver registered
		// with the global grpc resolver.
		grpc.WithResolvers(&fooResolverBuilder{}))
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

	dialWithFooResolver("resolver.foo.bar")
}

// Connecting to foo://resolver.foo.bar will fail unless we register
// a resolver that resolves foo://resolver.foo.bar to localhost:50505.
// Absence of a resolver for the foo scheme will result in a transport error.
// Here is a sample error
// err=rpc error: code = Unavailable desc = connection error: desc = "transport: Error while dialing dial tcp: lookup tcp////resolver.foo.bar: nodename nor servname provided, or not known"

type fooResolverBuilder struct{}

// ResolverBuilder implements the resolver.Builder interface.
// https://pkg.go.dev/google.golang.org/grpc/resolver?utm_source=godoc#Builder

func (f *fooResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	// Resolver implements the resolver.Resolver interface.
	// https://pkg.go.dev/google.golang.org/grpc/resolver?utm_source=godoc#Resolver
	r := &fooResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]resolver.Address{
			// resolve the foo service to localhost:50505
			"resolver.foo.bar": {
				{Addr: "localhost:50505"},
			},
		},
	}
	r.start()
	return r, nil
}

func (f *fooResolverBuilder) Scheme() string {
	return "foo"
}

// fooResolver is a custom resolver for the "foo" scheme.
// It implements the resolver.Resolver interface.
// https://pkg.go.dev/google.golang.org/grpc/resolver?utm_source=godoc#Resolver
type fooResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]resolver.Address
}

func (r *fooResolver) start() {
	// Sorry Copilot, you gave the deprecated method.
	// r.cc.NewAddress(r.addrsStore[r.target.Endpoint])

	// target.Endpoint is deprecated.
	// Even the example code in grpc-go uses target.Endpoint.
	//
	// Use URL.Path instead
	// URL.Path gives the path including the leading slash.
	endpoint := strings.TrimPrefix(r.target.URL.Path, "/")
	r.cc.UpdateState(resolver.State{Addresses: r.addrsStore[endpoint]})
}

func (r *fooResolver) ResolveNow(o resolver.ResolveNowOptions) {}
func (r *fooResolver) Close()                                  {}

// func init() {
// 	resolver.Register(&fooResolverBuilder{})
// }
