package resolver

import (
	"strings"

	"google.golang.org/grpc/resolver"
)

var ServerAddresses = []string{"localhost:50505"}

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
	// So trim the leading slash to get the address.
	address := strings.TrimPrefix(r.target.URL.Path, "/")
	r.cc.UpdateState(resolver.State{Addresses: r.addrsStore[address]})
}

func (r *fooResolver) ResolveNow(o resolver.ResolveNowOptions) {}
func (r *fooResolver) Close()                                  {}

type FooResolverBuilder struct{}

// ResolverBuilder implements the resolver.Builder interface.
// https://pkg.go.dev/google.golang.org/grpc/resolver?utm_source=godoc#Builder

func (f *FooResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	addresses := make([]resolver.Address, len(ServerAddresses))
	for i, address := range ServerAddresses {
		addresses[i] = resolver.Address{Addr: address}
	}

	// Resolver implements the resolver.Resolver interface.
	// https://pkg.go.dev/google.golang.org/grpc/resolver?utm_source=godoc#Resolver
	r := &fooResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]resolver.Address{
			"resolver.foo.bar": addresses,
		},
	}
	r.start()
	return r, nil
}

func (f *FooResolverBuilder) Scheme() string {
	return "foo"
}
