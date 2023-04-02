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

	myShelf := pb.Shelf{
		BooksToRead: []*pb.Book{&lostSymbol, &origin},
		BooksRead:   []*pb.Book{&steveJobs},
	}

	for index, book := range myShelf.BooksToRead {
		fmt.Printf("%v: title: %v | year: %v\n", index, book.Title, book.Year)
	}

	bookList1 := &pb.BookList{Books: []*pb.Book{
		{Title: "Book1", Year: 2021, Price: 29.99, IsReleased: true, Genre: pb.BookGenre_FICTION},
		{Title: "Book2", Year: 2022, Price: 19.99, IsReleased: false, Genre: pb.BookGenre_MEMOIR},
	}}

	bookList2 := &pb.BookList{Books: []*pb.Book{
		{Title: "Book3", Year: 2023, Price: 9.99, IsReleased: true, Genre: pb.BookGenre_THRILLER},
		{Title: "Book4", Year: 2024, Price: 14.99, IsReleased: false, Genre: pb.BookGenre_MEMOIR},
	}}

	booksByAuthorResp := &pb.BooksByAuthorResponse{Books: map[string]*pb.BookList{
		"Author1": bookList1,
		"Author2": bookList2,
	}}

	for author, booklist := range booksByAuthorResp.Books {
		fmt.Printf("%v -> %v", author, booklist)
	}
}
