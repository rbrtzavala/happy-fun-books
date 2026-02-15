package main

import (
	"books"
	"fmt"
)

func main() {
	catalog := books.GetCatalog() 
	for _, book := range catalog.GetAllBooks() {
		fmt.Println("Book in stock:")
		fmt.Println(book)
	}
}