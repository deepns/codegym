package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/deepns/codegym/go/learning/grpc/examples/reminders/reminders"
	"google.golang.org/grpc"
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
	return &pb.CreateReminderResponse{Id: int32(len(s.reminders))}, nil
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

func startServer(port int) {
	// create a listener on TCP port
	log.Printf("Starting server on port %v", port)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))
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

func main() {
	port := flag.Int("port", 50505, "port to connect to")
	flag.Parse()

	startServer(*port)
}
