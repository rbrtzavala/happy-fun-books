package books_test

import (
	"books"
	"cmp"
	"slices"
	"testing"
)


func TestGetAllBooks_ReturnsAllBooks(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	want := []books.Book {
		{
			Title: "In the Company of Cheerful Ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
			ID: "abc",
		},
		{
			Title: "White Heat",
			Author: "Dominic Sandbrook",
			Copies: 2,
			ID: "xyz",
		},
	}
	got := catalog.GetAllBooks()
	slices.SortFunc(got, func(a, b books.Book) int {
		return cmp.Compare(a.Author, b.Author)
	})
	if !slices.Equal(want, got) {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func TestGetBook_FindsBookInCatalogById(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	want := books.Book {
		ID: "abc",
		Title: "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
	}
	got, ok := catalog.GetBook("abc")
	if !ok {
		t.Fatalf("book not found :/")
	}
	if want != got {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func TestGotBook_ReturnsFalseWhenBookNotFount(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := catalog.GetBook("nonexistent ID")
	if ok {
		t.Fatalf("want false for nonexistent ID, got true")
	}
}

func TestAddBook_AddsGivenBookToCatalog(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := catalog.GetBook("123")
	if ok {
		t.Fatalf("book with already exists in catalog.")
	}
	catalog.AddBook(books.Book{
		ID: "123",
		Title: "The Prize of All Oceans",
		Author: "Glyn Williams",
		Copies: 3,
	})
	_, ok = catalog.GetBook("123")
	if !ok {
		t.Fatalf("added book not found 🤷‍♂️")
	}
}

func getTestCatalog() books.Catalog {
	return books.Catalog{
		"abc": {
			Title: "In the Company of Cheerful Ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
			ID: "abc",
	},
		"xyz": {
			Title: "White Heat",
			Author: "Dominic Sandbrook",
			Copies: 2,
			ID: "xyz",
		},
	}
}

func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	t.Parallel()
	input := books.Book {
		Title: "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}
	want := "Sea Room by Adam Nicolson (copies: 2)"
	got := input.String()
	if want != got {
		t.Fatalf("want %q got %q", want, got)
	}
}
