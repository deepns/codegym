package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/deepns/codegym/go/learning/grpc/examples/reminders/reminders"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateReminder(client pb.ReminderServiceClient, reminder string, when time.Time) {
	request := pb.Reminder{
		What: reminder,
		When: timestamppb.New(when),
		Type: pb.ReminderType_PUSH,
	}

	resp, err := client.CreateReminder(context.Background(), &request)
	if err != nil {
		log.Fatalf("Failed to create reminder. err=%v", err)
	}
	log.Printf("Created reminder. id=%v", resp.Id)
}

func ListReminders(client pb.ReminderServiceClient) {
	resp, err := client.GetReminders(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Failed to list reminders. err=%v", err)
	}

	for id, reminder := range resp.Reminders {
		log.Printf("Reminder{%v}: %v", id+1, reminder)
	}
}

func main() {
	addr := flag.String("addr", "localhost:50505", "address to connect to")
	create := flag.Bool("create", false, "create some reminders")
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server. err=%v", err)
	}
	defer conn.Close()

	client := pb.NewReminderServiceClient(conn)
	if *create {
		CreateReminder(client, "commit changes", time.Now().AddDate(0, 0, 1))
		CreateReminder(client, "send pr", time.Now().AddDate(0, 0, 2))
	}

	ListReminders(client)
}
