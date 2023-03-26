package main

import (
	"fmt"

	pb "github.com/deepns/codegym/go/learning/grpc/protobuf/books"
)

func main() {
	steveJobs := pb.Book{
		Title:      "Steve Jobs",
		Year:       2011,
		Price:      14.99,
		Genre:      pb.BookGenre_MEMOIR,
		IsReleased: true,
	}

	origin := pb.Book{
		Title:      "Origin",
		Year:       2017,
		Price:      16.50,
		Genre:      pb.BookGenre_THRILLER,
		IsReleased: true,
	}

	lostSymbol := pb.Book{
		Title:      "The Lost Symbol",
		Year:       2009,
		Price:      7.50,
		Genre:      pb.BookGenre_THRILLER,
		IsReleased: true,
	}

	myShelf := pb.Bookshelf{Books: []*pb.Book{&steveJobs, &origin}}
	for index, book := range myShelf.Books {
		fmt.Printf("%v: title: %v | year: %v\n", index, book.Title, book.Year)
	}

	// TODO
	// How to serialize the books?

	danBrownBooks := pb.Bookshelf{Books: []*pb.Book{&origin, &lostSymbol}}

	// The naming (or probably the composition) looks little misplaced..
	// That's okay for this example.
	bookMap := pb.BooksByAuthor{}
	bookMap.Books["Walter Issacson"] = &pb.Bookshelf{Books: []*pb.Book{&steveJobs}}
	bookMap.Books["Dan Brown"] = &danBrownBooks
}
