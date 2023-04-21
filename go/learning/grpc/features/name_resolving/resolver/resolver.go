package resolver

import (
	"strings"

	"google.golang.org/grpc/resolver"
)

type FooResolverBuilder struct{}

// ResolverBuilder implements the resolver.Builder interface.
// https://pkg.go.dev/google.golang.org/grpc/resolver?utm_source=godoc#Builder

func (f *FooResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
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

func (f *FooResolverBuilder) Scheme() string {
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
	// So trim the leading slash to get the address.
	address := strings.TrimPrefix(r.target.URL.Path, "/")
	r.cc.UpdateState(resolver.State{Addresses: r.addrsStore[address]})
}

func (r *fooResolver) ResolveNow(o resolver.ResolveNowOptions) {}
func (r *fooResolver) Close()                                  {}
