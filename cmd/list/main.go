package main

import (
	"books"
	"fmt"
)

func main() {
	fmt.Println("Books in stock:")
	for _, book := range books.GetAllBooks() {
		
		fmt.Println(books.BookToString(book))
	}
}