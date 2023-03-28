// Exploring the serialization of protobuf defined types
// in both binary and json format
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	pb "github.com/deepns/codegym/go/learning/grpc/protobuf/books"
	"google.golang.org/protobuf/proto"
)

func main() {
	shelf := &pb.Shelf{
		BooksToRead: []*pb.Book{
			{
				Title:      "Sapiens: A Brief History of Humankind",
				Year:       2015,
				Price:      12.99,
				IsReleased: true,
				Genre:      pb.BookGenre_MEMOIR,
			},
			{
				Title:      "The Water Dancer",
				Year:       2019,
				Price:      11.99,
				IsReleased: true,
				Genre:      pb.BookGenre_FICTION,
			},
		},
		BooksRead: []*pb.Book{
			{
				Title:      "The Underground Railroad",
				Year:       2016,
				Price:      9.99,
				IsReleased: true,
				Genre:      pb.BookGenre_FICTION,
			},
			{
				Title:      "The Power of Now: A Guide to Spiritual Enlightenment",
				Year:       1997,
				Price:      7.99,
				IsReleased: true,
				Genre:      pb.BookGenre_MEMOIR,
			},
		},
	}

	// Marshal the "shelf" instance into binary format
	shelfData, err := proto.Marshal(shelf)
	if err != nil {
		log.Fatalln("Error marshaling shelf data", err)
	}

	// Write the binary data to a file
	// Just want to see how the eventual data look like in either format
	err = ioutil.WriteFile("mybookshelf.pb", shelfData, 0644)
	if err != nil {
		log.Fatalln("Error writing binary data to file:", err)
	}

	// Marshal the "shelf" instance into json format
	// Since json encoding is defined in the struct itself, protobuf struct
	// can be readily marshaled to json.
	shelfDataJson, err := json.Marshal(shelf)
	if err != nil {
		log.Fatalln("Error marshaling shelf data to json", err)
	}

	err = ioutil.WriteFile("mybookshelf.json", shelfDataJson, 0644)
	if err != nil {
		log.Fatalln("Error writing json data to file:", err)
	}

	log.Println("Bytes written in pb format:", len(shelfData))
	log.Println("Bytes written in json format:", len(shelfDataJson))
}
