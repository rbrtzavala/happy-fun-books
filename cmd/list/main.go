package main

import (
	"books"
	"fmt"
)

func main() {
	book := books.Book {
		Title: "Engineering in Plain Sight",
		Author: "Grady Hollhouse",
		Copies: 142,
	}
    fmt.Println("Books in stock:")
    fmt.Println(books.BookToString(book))
}