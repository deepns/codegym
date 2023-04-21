package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/deepns/codegym/go/learning/grpc/examples/reminders/reminders"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type reminderServer struct {
	pb.UnimplementedReminderServiceServer
	// Just a simple in-memory storage for the reminders
	reminders []*pb.Reminder
}

func (s *reminderServer) CreateReminder(ctx context.Context,
	req *pb.Reminder) (*pb.CreateReminderResponse, error) {
	log.Printf("CreateReminder: %v", req)
	s.reminders = append(s.reminders, &pb.Reminder{What: req.What, When: req.When, Type: req.Type})
	return &pb.CreateReminderResponse{Id: int32(len(s.reminders)), Success: true}, nil
}

func (s *reminderServer) GetReminders(ctx context.Context, req *pb.Empty) (*pb.GetRemindersResponse, error) {
	log.Printf("GetReminders: %v", s.reminders)
	return &pb.GetRemindersResponse{Reminders: s.reminders}, nil
}

func (s *reminderServer) GetReminder(ctx context.Context, req *pb.GetReminderRequest) (*pb.Reminder, error) {
	log.Printf("GetReminder: %v", req)
	if req.Id < 1 || int(req.Id) > len(s.reminders) {
		return nil, fmt.Errorf("invalid reminder id %v", req.Id)
	}
	return s.reminders[req.Id-1], nil
}

func startGrpcServer(port int) {
	// create a listener on TCP port
	log.Printf("Starting server on port %v", port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to listen on %v. err=%v", port, err)
	}

	// not using tls. so just an empty option would do.
	var options []grpc.ServerOption
	grpcServer := grpc.NewServer(options...)

	// register the server with the grpc service
	pb.RegisterReminderServiceServer(grpcServer, &reminderServer{})

	// register reflection service on gRPC server.
	reflection.Register(grpcServer)

	// serve on the created listener
	grpcServer.Serve(lis)
}

// startGatewayServer makes a connection to the grpc server
// and starts the gRPC gateway server
func startGatewayServer(grpcServerPort, grpcGatewayPort int) {
	// create a gRPC server client connection
	gmux := runtime.NewServeMux()

	// A Note on the various RegisterReminderServiceHandler* methods
	// RegisterReminderServiceHandler is just a convenience wrapper around RegisterReminderServiceHandlerClient
	// RegisterReminderServiceHandlerFromEndpoint makes this even simpler by taking care of the dialing
	// and context creation and cancellation.
	// RegisterReminderServiceHandlerServer registers the http handlers for the given server.
	// Doc says this registration option will cause many features of gRPC library to stop working,
	// and recommends using RegisterReminderServiceHandlerFromEndpoint instead.

	// RegisterReminderServiceHandlerFromEndpoint internally creates a connection
	// and registers the connection to proxy requests to the gRPC server.
	err := pb.RegisterReminderServiceHandlerFromEndpoint(context.Background(),
		gmux,
		fmt.Sprintf("0.0.0.0:%v", grpcServerPort),
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})

	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%v", grpcGatewayPort),
		Handler: gmux,
	}

	log.Printf("Starting gateway server on port %v", grpcGatewayPort)
	gwServer.ListenAndServe()
}

func main() {
	port := flag.Int("grpc-port", 50505, "port to connect to")
	grpcGatewayPort := flag.Int("grpc-gateway-port", 8080, "port to connect to")
	flag.Parse()

	// grpc server blocks on the Serve() method. So, if we want to start the gateway server
	// in the same process, we need to start the grpc server in a goroutine.
	go startGrpcServer(*port)
	startGatewayServer(*port, *grpcGatewayPort)
}
