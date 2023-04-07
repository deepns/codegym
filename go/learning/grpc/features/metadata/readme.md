# GRPC Metadata

## What it is?

- list of key-value pairs carrying information about a particular RPC call
- keys are strings and values can be strings or binary data
- gRPC defined or user-defined

## Where is it used?

- Authentication using SSL/TLS, ATLS or OAuth2 tokens carry the authentication information in the metadata
- for e.g. when authentication using OAuth2 token, server side intercepts the call, extracts the token from the metadata and validates it before the actual handler for the RPC call is invoked

## Key Naming

- keys starting with `grpc-` are reserved for grpc
- keys carrying binary values end in `-bin`
- case insensitive
- allowed characters - ASCII letters, digits, and special characters `-, _, .`

## Constraints

- Access to metadata is language dependent.
- Go supports it. Not sure about the other languages

## Language Specifics

- Metadata package in [grpc-go](https://pkg.go.dev/google.golang.org/grpc@v1.52.0/metadata#MD)
- MD is a mapping of key to a list of values. Create metadata using the convenience functions defined in `metadata` package

```go
timestampMD := metadata.New(map[string]string{"timestamp": time.Now().Format(time.StampNano)})
token := metadata.Pairs("token", "1234567890")
```

- reading a metadata from a client connection

```go
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "missing metadata")
	}

	for k, v := range md {
		log.Printf("metadata: %v=%v", k, strings.Join(v, ","))
	}
```
