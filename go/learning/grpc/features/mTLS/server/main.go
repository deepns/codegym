package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	"github.com/deepns/codegym/go/learning/grpc/features/sslcerts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type echoServer struct {
	pb.UnimplementedEchoServiceServer
}

func (s *echoServer) UnaryEcho(_ context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("UnaryEcho: %v", req.Message)
	return &pb.EchoResponse{Message: req.Message}, nil
}

func getTLSConfig() *tls.Config {
	serverCert, err := tls.LoadX509KeyPair(
		sslcerts.Path("server_cert.pem"), sslcerts.Path("server_key.pem"))
	if err != nil {
		log.Fatalf("failed to load server key pair: %v", err)
	}

	caCert, err := ioutil.ReadFile(sslcerts.Path("client_ca_cert.pem"))
	if err != nil {
		log.Fatalf("failed to load CA certificate: %v", err)
	}

	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatalf("failed to append CA certificate")
	}

	return &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}
}

func main() {
	port := flag.Int("port", 50505, "port to listen")
	flag.Parse()

	tlsConfig := getTLSConfig()
	s := grpc.NewServer(grpc.Creds(credentials.NewTLS(tlsConfig)))
	pb.RegisterEchoServiceServer(s, &echoServer{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on port %d", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
