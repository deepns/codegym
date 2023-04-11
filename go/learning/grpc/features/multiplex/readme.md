# gRPC multiplexing

Didn't know that server can mulitiplex RPCs from different services.

## Server side

In this example, EchoService exposes these RPCs.

```proto
service EchoService {
    // An Unary RPC. Send the request and get the response
    rpc UnaryEcho(EchoRequest) returns (EchoResponse) {}

    // Streams the response from the client side
    rpc ClientSideStreamEcho(stream EchoRequest) returns (EchoResponse) {}

    // Streams the response from the server side
    rpc ServerSideStreamEcho(EchoRequestWithCount) returns (stream EchoResponse) {}

    // A Bidirectional streaming RPC
    //
    // Accepts a stream of messages and echoes it back as it receives
    rpc BidrectionalStreamEcho(stream EchoRequest) returns (stream EchoResponse) {}
}
```

and, HelloWorld exposes these RPCs

```proto
service HelloService {
    rpc SayHello(HelloRequest) returns (HelloResponse);
}
```

A server can be configured to serve both the services, as long as the type implements the required RPCs. Create the server types, register with the equivalent Register... function before serving.

```go
	s := grpc.NewServer()
	pbec.RegisterEchoServiceServer(s, &echoServer{})
	pbhw.RegisterHelloServiceServer(s, &helloServer{})

	log.Printf("Listening on port %d", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
```

## Client side

- How to share a connection between two stubs?
- By creating a new client with the dialed connection

```go
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server. err=%v", err)
	}
	defer conn.Close()

	echoClient := pbec.NewEchoServiceClient(conn)
	callUnaryEcho(echoClient, "Boo!!!")

	helloClient := pbhw.NewHelloServiceClient(conn)
	callSayHello(helloClient, "Steve")
```

- The service running in this example serves both Echo service and Hello service. so we are able to use the same connection to send rpc of different services
