package main

import "fmt"

// type Author struct {
// 	FirstName string
// 	LastName  string
// }

type Book struct {
	ID     int
	Title  string
	Year   int //string
	Author string
	// BookAuthor Author
	IsIssued bool
}

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

type Library struct {
	Books   []*Book
	Readers []*Reader

	lastBookID   int
	lastReaderID int
}

func (b Book) SeeInfoBook() {
	fmt.Printf("Название: %s, Год: %v, Автор: %v. Книга в наличии: %v\n", b.Title, b.Year /*b.BookAuthor.FirstName, b.BookAuthor.LastName,*/, b.Author, b.IsIssued)
}

func (b *Book) IssueBook() {
	if b.IsIssued {
		fmt.Printf("\nКнига %s уже кому-то выдана\n", b.Title)
		return
	}
	b.IsIssued = true
	fmt.Printf("\nКнига %s была выдана\n", b.Title)
}

func (lib *Library) AddReader(firstname, lastname string) *Reader {
	lib.lastReaderID++

	newReader := &Reader{
		ID:        lib.lastReaderID,
		FirstName: firstname,
		LastName:  lastname,
		IsActive:  true,
	}

	lib.Readers = append(lib.Readers, newReader)

	fmt.Printf("Зарегистрировался читатель: %v %v\n", newReader.FirstName, newReader.LastName)
	return newReader
}

func (lib *Library) AddBook(title, author string, year int) *Book {
	lib.lastBookID++

	newBook := &Book{
		ID:       lib.lastBookID,
		Title:    title,
		Author:   author,
		Year:     year,
		IsIssued: false,
	}

	lib.Books = append(lib.Books, newBook)

	fmt.Printf("Добавлена новая книга: %v\n", newBook)
	return newBook
}

func (lib *Library) FindBookById(id int) (*Book, error) {
	for i := 0; i < len(lib.Books); i++ {
		if id == lib.Books[i].ID {
			return lib.Books[i], nil
		}
	}
	return nil, fmt.Errorf("Книга с id %d не найдена\n", id)
}

func (lib *Library) FindReaderById(id int) (*Reader, error) {
	for i := 0; i < len(lib.Books); i++ {
		if id == lib.Readers[i].ID {
			return lib.Readers[i], nil
		}
	}
	return nil, fmt.Errorf("Читатель с id %d не найден\n", id)
}

func (lib *Library) IssueBookToReader(bookID int, readerID int) error {
	_, errBook := lib.FindBookById(bookID)
	if errBook != nil {
		return errBook
	}
	_, errReader := lib.FindReaderById(bookID)
	if errReader != nil {
		return errReader
	}

	lib.Books[bookID].IssueBook()
	if lib.Books[bookID].IsIssued {
		return nil
	} else {
		return fmt.Errorf("Произошла ошибка\n")
	}
}

func main() {
	// books := []Book{}
	// author1 := Author{
	// 	FirstName: "fname1",
	// 	LastName:  "lname1",
	// }
	// author2 := Author{
	// 	FirstName: "fname2",
	// 	LastName:  "lname2",
	// }
	// author3 := Author{
	// 	FirstName: "fname3",
	// 	LastName:  "lname3",
	// }
	// book1 := Book{
	// 	Title:      "title1",
	// 	Year:       "1111",
	// 	BookAuthor: author1,
	// }
	// book2 := Book{
	// 	Title:      "title2",
	// 	Year:       "2222",
	// 	BookAuthor: author2,
	// }
	// book3 := Book{
	// 	Title:      "title3",
	// 	Year:       "3333",
	// 	BookAuthor: author3,
	// }
	// books = append(books, book1, book2, book3)
	// for i := 0; i < len(books); i++ {
	// 	seeInfoBook(books[i])
	// }

	// slice := []Notifier{}

	// email := EmailNotifier{"www@gmail.com"}
	// phoneNum := SMSNotifier{"+79998887766"}

	// slice = append(slice, email, phoneNum)

	// for i := 0; i < len(slice); i++ {
	// 	slice[i].Notify("Ваша книга просрочена")
	// }

	myLibrary := &Library{}

	myLibrary.AddReader("Тамара", "Коляда")
	myLibrary.AddReader("Давид", "Хубаев")

	myLibrary.AddBook("Я чут-чут не книжный червь", "Т. Коляда", 2027)
	myLibrary.AddBook("Мифы древней Греции", "Греки Древние", 1990)

	fmt.Println("---Тестируем выдачу книг---")
	err := myLibrary.IssueBookToReader(1, 1)
	if err != nil {
		fmt.Printf("Ошибка выдачи: %v\n", err)
	}

	book, _ := myLibrary.FindBookById(1)
	if book != nil {
		fmt.Printf("Статус книги после выдачи: %v\n", book)
	}

	err2 := myLibrary.IssueBookToReader(99, 1)
	if err2 != nil {
		fmt.Printf("Ожидается ошибка	: %v\n", err2)
	}
}
