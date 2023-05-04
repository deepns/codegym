package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	"google.golang.org/grpc"
)

type echoServer struct {
	pb.UnimplementedEchoServiceServer
	id int
}

func (s *echoServer) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("ServerPort:%d UnaryEcho: %v", s.id, req.Message)
	return &pb.EchoResponse{Message: req.Message}, nil
}

func main() {
	ports := flag.String("ports", "50505", "comma separated ports to listen to")
	flag.Parse()

	for _, port := range strings.Split(*ports, ",") {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		id, _ := strconv.Atoi(port)
		pb.RegisterEchoServiceServer(s, &echoServer{id: id})

		log.Printf("Listening on port %v", port)
		go s.Serve(lis)
	}

	// Wait until Ctrl+C
	select {}

}
