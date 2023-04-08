# gRPC Interceptor

## Notes

- What interceptor does? Obviously, it intercepts. What it intercepts? - intercepts the execution of each RPC call
- can be configured independently on the client side and server side
- types of interceptors
  - Unary Interceptor
    - For Client
    - For Server
  - Stream Interceptor
    - For Client
    - For Server
- [UnaryClientInterceptor](https://pkg.go.dev/google.golang.org/grpc#UnaryClientInterceptor)
  - how is it configured? as a DialOption using `WithUnaryInterceptor(func)` or `WithChainUnaryInterceptor(func)`, when creating a ClientConn
  - what is the function signature?
    - `func(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, invoker UnaryInvoker, opts ...CallOption) error`
  - stages
    - pre-processing
    - invoking the actual RPC
    - post-processing