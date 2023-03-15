package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/route_guide/routeguide"
)

// Shows a sample invocation of a simple rpc request to the RouteGuide server
func GetFeatureSample(client pb.RouteGuideClient) {
	p1 := pb.Point{Latitude: 404318328, Longitude: -740835638}
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	feature, err := client.GetFeature(context, &p1)
	if err != nil {
		log.Fatalf("failed to get feature for %v, err=%c", p1, err)
	}
	log.Printf("Received %v for point %v:%v", feature, p1.Latitude, p1.Longitude)
}

func main() {
	var serverAddr = flag.String("addr", "localhost:30303", "host:port of the server addr to connect to")
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("connection to %v failed due to %v", serverAddr, err)
	}
	defer conn.Close()
	client := pb.NewRouteGuideClient(conn)

	GetFeatureSample(client)
}
