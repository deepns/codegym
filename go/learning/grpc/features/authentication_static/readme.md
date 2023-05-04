# Using a simple usernmae/password authentication in gRPC

Using static username and password for authentication is not ideally recommended for gRPC services. Just wanted to try it out to see how we can configure one.

The simplest way is to make use of metadata. We can pass in username and password strings in the metadata section from the client side. Server can intercept the rpc call using appropriate interceptor, extract the metadata and do the necessary authentication. Metadata field names must match on server and client side though.

We can pass the username and password through metadata in two ways

## Explicitly adding them to the outgoing context

This is little clumsy though, using PerRPCCredentials is the preferred way.

```go
ctx := metadata.NewOutgoingContext(context.Background(),
    metadata.Pairs("username", "guest", "password", "guest123"))
```

[clientv1](clientv1/main.go) is an example of this usage.

## Using [PerRPCCredentials](https://pkg.go.dev/google.golang.org/grpc@v1.54.0/credentials#PerRPCCredentials) interface

- Define the interface that returns the username and password in response to GetRequestMetadata

```go
type myCredential struct{}

func (c myCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"username": "admin",
		"password": "secret",
	}, nil
}

func (c myCredential) RequireTransportSecurity() bool { return true }
```

- And then dial to the server with `grpc.WithPerRPCCredentials` dial option.

```go
	conn, err := grpc.Dial(*addr,
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(myCredential{}))
```

[clientv2](clientv2/main.go) is an example of using PerRPCCredentials for username/password based authentication
