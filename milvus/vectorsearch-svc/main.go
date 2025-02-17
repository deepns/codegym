package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/deepns/codegym/milvus/vectorsearch-svc/proto"
	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	"google.golang.org/grpc"
)

const (
	milvusAddr = "localhost:19530"
	dimension  = 128 // Adjust based on your embeddings
	collection = "vector_collection"
)

type vectorServiceServer struct {
	pb.UnimplementedVectorServiceServer
	milvusClient client.Client
}

func (s *vectorServiceServer) InsertVector(ctx context.Context, req *pb.InsertVectorRequest) (*pb.InsertVectorResponse, error) {
	id := req.GetId()
	embedding := req.GetEmbedding()

	if len(embedding) != dimension {
		return nil, fmt.Errorf("embedding dimension mismatch")
	}

	// Convert to Milvus format
	data := entity.NewColumnFloatVector("embedding", dimension, [][]float32{embedding})
	ids := entity.NewColumnInt64("id", []int64{id})

	_, err := s.milvusClient.Insert(ctx, collection, "", ids, data)
	if err != nil {
		return nil, err
	}

	return &pb.InsertVectorResponse{Success: true}, nil
}

func (s *vectorServiceServer) SearchVector(ctx context.Context, req *pb.SearchVectorRequest) (*pb.SearchVectorResponse, error) {
	// embedding := req.GetEmbedding()
	// topK := req.GetTopK()

	// searchParams, _ := entity.NewIndexFlatSearchParam()

	// results, err := s.milvusClient.Search(ctx, collection, []string{"id"}, []entity.Vector{
	// 	entity.FloatVector(dimension, embedding)}, "", topK, searchParams)
	// if err != nil {
	// 	return nil, err
	// }

	var response pb.SearchVectorResponse

	return &response, nil
}

func newServer() (*vectorServiceServer, error) {
	milvusAddr := "localhost:19530"
	milvusClient, err := client.NewClient(context.Background(), client.Config{
		Address: milvusAddr,
	})
	if err != nil {
		return nil, err
	}

	return &vectorServiceServer{milvusClient: milvusClient}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	server, err := newServer()
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}
	pb.RegisterVectorServiceServer(grpcServer, server)

	log.Println("gRPC server running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
