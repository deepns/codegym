# gRPC reflection

## What reflection does?

- client has to know
  - what method does a server export?
  - Is the exported method Unary or Streaming?
  - What are the types of input and output arguments?
- reflection provides a way to get those information from server and interact (invoking RPCs) with it without having the stub information precompiled into it.

## How to enable reflection in server?

Import [reflection package](https://pkg.go.dev/google.golang.org/grpc/reflection) and register the service with the reflection using `reflection.Register`.

## Where is it used?

Primarily in command line debugging tools for a gRPC server. Tool take the payload in human readable (json, text) format, turn it into binary (proto) format and send it over the wire. Server responds in binary format (proto), and the tool converts it back into human readable (text, json)

## grpcurl

- A curl like tool for gRPC
- Install with `brew install grpcurl` or `arch -arm64 brew install grpcurl` for M1 Macs
- Sample invocations of grpcurl
  - `grpcurl grpc.server.com:443 my.custom.server.Service/Method` - for tls enabled server
  - `grpcurl -plaintext grpc.server.com:80 my.custom.server.Service/Method` - for insecure connections
- Enabled reflection on my test server and ran grpcurl

```text
✗ grpcurl -plaintext localhost:50505 list
EchoService
HelloService
grpc.reflection.v1alpha.ServerReflection
```

- Describing a service

```text
✗ grpcurl -plaintext localhost:50505 describe HelloService
HelloService is a service:
service HelloService {
  rpc SayHello ( .HelloRequest ) returns ( .HelloResponse );
}
```

- Describing a method.

```text
✗ grpcurl -plaintext localhost:50505 describe HelloService.SayHello
HelloService.SayHello is a method:
rpc SayHello ( .HelloRequest ) returns ( .HelloResponse );
```

- Describing a message..use the message name as it appears in the method description

```text
✗ grpcurl -plaintext localhost:50505 describe .HelloRequest        
HelloRequest is a message:
message HelloRequest {
  string name = 1;
  uint32 count = 2;
}
```

- Interacting with the server from `grpc-go/examples/features/reflections`...can see that services are defined in `<package-name>.<service-name>` format.

```text
✗ grpcurl -plaintext localhost:50506 list                          
grpc.examples.echo.Echo
grpc.reflection.v1alpha.ServerReflection
helloworld.Greeter
```

- And, so are the messages `<package-name>.<message-name>`

```text
✗ grpcurl -plaintext localhost:50506 describe helloworld.Greeter
helloworld.Greeter is a service:
service Greeter {
  rpc SayHello ( .helloworld.HelloRequest ) returns ( .helloworld.HelloReply );
  rpc SayHelloAgain ( .helloworld.HelloRequest ) returns ( .helloworld.HelloReply );
}
```

- Invoking a rpc..pass the service data in option `-d`

```text
✗ grpcurl -plaintext -d '{"name": "Foobar"}' localhost:50506 helloworld.Greeter/SayHello
{
  "message": "Hello Foobar"
}

✗ grpcurl -plaintext -d '{"name": "Foobar", "count": 2}' localhost:50505 HelloService.SayHello
{
  "helloMsg": "Hello, Foobar\nHello, Foobar\n"
}
```

- Lot more functionalities available in grpcurl.

## Notes

- [reflection protocol](https://github.com/grpc/grpc/blob/master/doc/server-reflection.md)
- [grpcurl on github](https://github.com/fullstorydev/grpcurl)
