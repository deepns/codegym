package main

import (
	"log"

	pb "github.com/deepns/codegym/go/learning/grpc/helloworld/helloworld"
	"github.com/golang/protobuf/proto"
)

func main() {
	request := pb.HelloRequest{Name: "FooBar", Count: 2}
	out, err := proto.Marshal(&request)
	// packs it this way..type identifier, length, bytes...
	// 10 for string, length 6
	// 16 for uint32
	// [10 6 70 111 111 66 97 114 16 2]
	if err != nil {
		log.Fatalf("Failed to marshal the request. err=%v", err)
	}
	log.Printf("Marshalling {%v:%v} %v", request.Name, request.Count, out)

	request.Name = "JohnDoe"
	out, err = proto.Marshal(&request)
	if err != nil {
		log.Fatalf("Failed to marshal the request. err=%v", err)
	}
	log.Printf("Marshalling {%v:%v} %v", request.Name, request.Count, out)
}
