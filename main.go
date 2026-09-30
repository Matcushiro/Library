package main

import "fmt"

type Author struct {
	FirstName string
	LastName  string
}

type Book struct {
	Title      string
	Year       string
	BookAuthor Author
}

func seeInfoBook(b Book) {
	fmt.Printf("Название: %s, Год: %s, Автор: %v %v\n", b.Title, b.Year, b.BookAuthor.FirstName, b.BookAuthor.LastName)
}

func main() {
	books := []Book{}
	author1 := Author{
		FirstName: "fname1",
		LastName:  "lname1",
	}
	author2 := Author{
		FirstName: "fname2",
		LastName:  "lname2",
	}
	author3 := Author{
		FirstName: "fname3",
		LastName:  "lname3",
	}
	book1 := Book{
		Title:      "title1",
		Year:       "1111",
		BookAuthor: author1,
	}
	book2 := Book{
		Title:      "title2",
		Year:       "2222",
		BookAuthor: author2,
	}
	book3 := Book{
		Title:      "title3",
		Year:       "3333",
		BookAuthor: author3,
	}
	books = append(books, book1, book2, book3)
	for i := 0; i < len(books); i++ {
		seeInfoBook(books[i])
	}
}
