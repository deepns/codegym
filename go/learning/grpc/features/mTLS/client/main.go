package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	"github.com/deepns/codegym/go/learning/grpc/features/sslcerts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func getCredentials() credentials.TransportCredentials {
	cert, err := tls.LoadX509KeyPair(
		sslcerts.Path("client_cert.pem"), sslcerts.Path("client_key.pem"))
	if err != nil {
		log.Fatalf("failed to load key pair: %v", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(sslcerts.Path("ca_cert.pem"))
	if err != nil {
		log.Fatalf("failed to read CA certificate: %v", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("failed to append CA certificate")
	}

	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "abc.test.example.com",
		RootCAs:      certPool,
	})
}

func main() {
	addr := flag.String("addr", "localhost:50505", "address to connect to")
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(getCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)
	resp, err := client.UnaryEcho(context.Background(), &pb.EchoRequest{Message: "hello from mtls client"})
	if err != nil {
		log.Fatalf("failed to call UnaryEcho: %v", err)
	}

	log.Printf("echo: %s", resp.Message)
}
