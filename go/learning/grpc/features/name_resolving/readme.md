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
  - where is it used? #TODO
- Resolvers are built using [ResolverBuilder](https://pkg.go.dev/google.golang.org/grpc/resolver?utm_source=godoc#Builder). Builder and [Resolver](https://pkg.go.dev/google.golang.org/grpc/resolver?utm_source=godoc#Resolver) are interfaces
- Need to go through the example resolver and passthrough resolver to learn more how they are built, how and where they are used
  