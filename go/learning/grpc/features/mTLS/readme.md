# Configuring gRPC client-server with mutual TLS

## Server side

- When configured just TLS, I used the convenience function [NewServerTLSFromFile](https://pkg.go.dev/google.golang.org/grpc@v1.53.0/credentials#NewServerTLSFromFile) to generate the transport credentials. This is just a wrapper around [credentials.NewTLS](https://github.com/grpc/grpc-go/blob/v1.53.0/credentials/tls.go#L142) that takes a [TLS config](https://pkg.go.dev/crypto/tls#Config) and returns [TransportCredentials](https://pkg.go.dev/google.golang.org/grpc@v1.53.0/credentials#TransportCredentials)

```go
	cert, err := credentials.NewServerTLSFromFile(
		sslcerts.Path("server_cert.pem"),
		sslcerts.Path("server_key.pem"))
	if err != nil {
		log.Fatalf("failed to load TLS cert: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(cert))
```

- To configure the server with mutual TLS, I need to provide the client CA certificates and also tell gRPC to verify the client cert with the provided client CAs. How do I provide?
- Create a new transport credential using [credentials.NewTLS](https://github.com/grpc/grpc-go/blob/v1.53.0/credentials/tls.go#L142) with a `tls.Config` configured holding all the necessary info
  - server certificate(**tls.Config.Certificates**)
  - client CA certs (**tls.Config.ClientCAs**)
  - authentication type (**tls.Config.ClientAuth**)

```go
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientCAs:    clientCAPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
    creds := credentials.NewTLS(tlsConfig)
```

- To load the server cert, do what NewServerTLSFromFile does. i.e. Load the certificate and key using [LoadX509KeyPair](https://pkg.go.dev/crypto/tls#LoadX509KeyPair)

```go
serverCert, err := tls.LoadX509KeyPair("server_cert.pem", "server_key.pem")
if err != nil {
    log.Fatalf("failed to load server cert: %v", err)
}
```

- To provide list of client CAs, create a cert pool and append every client CA to the cert pool. `AppendCertsFromPEM` is the preferred way to append to the cert pool.

```go
	// Load client CA certificate
	clientCACert, err := ioutil.ReadFile("client_ca.crt")
	if err != nil {
		log.Fatalf("Failed to load client CA certificate: %v", err)
	}
	clientCAPool := x509.NewCertPool()
	if ok := clientCAPool.AppendCertsFromPEM(clientCACert); !ok {
		log.Fatalf("Failed to append client CA certificate")
	}
```

- I see [certPooo.AddCert](https://pkg.go.dev/crypto/x509#CertPool.AddCert) as wel. But AddCert cannot be used to append the client CA to the pool because it expects a single DER-encoded X.509 certificate as input, whereas AppendCertsFromPEM is used to add one or more PEM-encoded certificates to the pool.

The AddCert method is typically used to add a single self-signed certificate to a pool, whereas in the case of mutual TLS, the CA may have issued multiple certificates that are needed to validate the client's certificate.

## Client side

When I configured TLS on the client side, I used the convenient function `NewClientTLSFromFile` to get the transport credentials.

```go
	creds, err := credentials.NewClientTLSFromFile(sslcerts.Path("ca_cert.pem"), "abc.test.example.com")
	if err != nil {
		log.Fatalf("failed to load TLS cert: %v", err)
	}
```

Like the server side, this also creates a new cert pool, include the root CA to validate server certificate and create a TL config with those parameters. If client certificate needs to be included for mutual TLS, I got to use **NewTLS** instead.

```go
// NewClientTLSFromFile constructs TLS credentials from the provided root
// certificate authority certificate file(s) to validate server connections. If
// certificates to establish the identity of the client need to be included in
// the credentials (eg: for mTLS), use NewTLS instead, where a complete
// tls.Config can be specified.
// serverNameOverride is for testing only. If set to a non empty string,
// it will override the virtual host name of authority (e.g. :authority header
// field) in requests.
func NewClientTLSFromFile(certFile, serverNameOverride string) (TransportCredentials, error) {
	b, err := os.ReadFile(certFile)
	if err != nil {
		return nil, err
	}
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		return nil, fmt.Errorf("credentials: failed to append certificates")
	}
	return NewTLS(&tls.Config{ServerName: serverNameOverride, RootCAs: cp}), nil
```

In addition to the **ServerName** and **RootCAs**, need to provide client certificate in **Certificates**.

```go
    cert, err := tls.LoadX509KeyPair(
		sslcerts.Path("client_cert.pem"), sslcerts.Path("client_key.pem"))
    //...
    //...
    return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "abc.test.example.com",
		RootCAs:      certPool,
	})
```
