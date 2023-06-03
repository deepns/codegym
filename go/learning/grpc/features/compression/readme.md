# Enabling compression on client and server side

To configure message compression, register a compressor using [`encoding.RegisterCompressor`](https://godoc.org/google.golang.org/grpc/encoding#RegisterCompressor). grpc-go provides [gzip](https://github.com/grpc/grpc-go/blob/v1.55.0/encoding/gzip/gzip.go) compression in its library. 

## gzip

package [gzip](https://github.com/grpc/grpc-go/blob/v1.55.0/encoding/gzip/gzip.go) implements and register a gzip compressor. gzip's `init()` registers the compressor using `encoding.RegisterCompressor`, which adds the compressor to `registeredCompressor` map. grpc server and client code gets the registered compressor through `GetCompressor` when invoking the registered RPCs.

```go
func init() {
	c := &compressor{}
	c.poolCompressor.New = func() interface{} {
		return &writer{Writer: gzip.NewWriter(io.Discard), pool: &c.poolCompressor}
	}
	encoding.RegisterCompressor(c)
}
```

## Configuring compression

- On the server side, not much to do except to register the compressor. Importing gzip package takes care of registering the gzip compressor through its init function.
- On the client side, we can use custom compressor or use gzip
  - To configure gzip compressor for all calls on the connection, use `grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name))`

    ```go
    conn, err := grpc.Dial(*addr,
            grpc.WithTransportCredentials(insecure.NewCredentials()),
            grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
    ```

  - To configure gzip compressor per call, use `grpc.UseCompressor(gzip.Name)` in the call options

    ```go
        resp, err := client.UnaryEcho(ctx, &pb.EchoRequest{Message: msg}, grpc.UseCompressor(gzip.Name))
    ```
