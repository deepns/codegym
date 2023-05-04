# Name resolving

## About name resolving in gRPC

- Reading about name resolving from the [doc](https://github.com/grpc/grpc/blob/master/doc/naming.md)
- Been running my servers at `localhost:port` all along, so name resolution wasn't needed. If I had use a fully qualified address (like a [URI](https://tools.ietf.org/html/rfc3986)), it would have resolved the address to ipaddress:port using DNS
- Support for different URI schemes are implementation dependent. Most gRPC implementations supports
  - **dns** - `dns:[//authority/]host[:port]`
  - **unix domain sockets** - `unix:path`, `unix://absolute_path`
  - **unix domain sockets in abstract namespace** - `unix-abstract:abstract_path`
- gRPC C-core implementation supports the following schemes in addition to the above
  - **ipv4** - `ipv4:address[:port][,address[:port],...]`
  - **ipv6** - `ipv6:address[:port][,address[:port],...]`

## name resolving in grpc-go

- Supports **dns, unix and passthrough** resolvers - [doc](https://pkg.go.dev/google.golang.org/grpc/internal/resolver#section-directories)
  - [passthrough](https://github.com/grpc/grpc-go/blob/v1.54.0/internal/resolver/passthrough/passthrough.go) sends the target name without the scheme back to gRPC as resolved address.
- Find some places where name resolving is used
  - [etcd client](https://github.com/etcd-io/etcd/blob/b27dec8b9487d0d9358ca4dd366563d1aab04a1e/client/v3/naming/resolver/resolver.go) supports grpc resolver for its name resolution.
- How it connects using the resolver? **Note**: resolver can be registered with the global register via init() and calling `resolver.Register` or registered for the specific dial using `grpc.WithResolvers`

```go
// Dial an RPC service using the etcd gRPC resolver and a gRPC Balancer:
// example:
func etcdDial(c *clientv3.Client, service string) (*grpc.ClientConn, error) {
    etcdResolver, err := resolver.NewBuilder(c);
    if err { return nil, err }
    return  grpc.Dial("etcd:///" + service, grpc.WithResolvers(etcdResolver))
}
```

- `grpc-go` says [grpc.WithResolvers](https://pkg.go.dev/google.golang.org/grpc#WithResolvers) experimental though

```go
func WithResolvers(rs ...resolver.Builder) DialOption {
	return newFuncDialOption(func(o *dialOptions) {
		o.resolvers = append(o.resolvers, rs...)
	})
}
```

- Resolvers are built using [ResolverBuilder](https://pkg.go.dev/google.golang.org/grpc/resolver?utm_source=godoc#Builder). Builder and [Resolver](https://pkg.go.dev/google.golang.org/grpc/resolver?utm_source=godoc#Resolver) are interfaces
  - The typical workflow is to implement the Builder interface that builds a struct implementing Resolver interface. An example resolver.

    ```go
    type fooResolver struct {
        target     resolver.Target
        cc         resolver.ClientConn
        addrsStore map[string][]resolver.Address
    }
    ```

  - Update the ClientConn with [resolver.State](https://pkg.go.dev/google.golang.org/grpc@v1.52.0/resolver#State) holding the address given by the builder

    ```go
    func (r *fooResolver) start() {
        // target.Endpoint is deprecated. Use URL.Path instead.
        // URL.Path gives the path including the leading slash.
        // So trim the leading slash to get the address.
        address := strings.TrimPrefix(r.target.URL.Path, "/")
        r.cc.UpdateState(resolver.State{Addresses: r.addrsStore[address]})
    }
    ```
