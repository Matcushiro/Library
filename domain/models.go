package domain

import "fmt"

type Book struct {
	ID       int
	Title    string
	Author   Author
	IsIssued bool
	ReaderID int
}

type Reader struct {
	ID       int
	Name     string
	IsActive bool
}

type Author struct {
	Name    string
	Surname string
}

//Методы книг

func (b *Book) String() string {
	id := 0
	if b.IsIssued {
		id = b.ReaderID
	}
	return fmt.Sprintf("%d: %s от автора %s %s. Выдана: %v. ID пользователя: %d\n", b.ID, b.Title, b.Author.Name, b.Author.Surname, b.IsIssued, id)
}

func (b *Book) IssueBook(r Reader) error {
	if b.IsIssued {
		return fmt.Errorf("Книга %s от автора %s %s уже выдана другому читателю \n", b.Title, b.Author.Name, b.Author.Surname)
	}
	if !r.IsActive {
		return fmt.Errorf("Читатель %s отключен и не может получить книгу\n", r.Name)
	}
	b.ReaderID = r.ID
	b.IsIssued = true
	return nil
}

func (b *Book) ReturnBook() error {
	if !b.IsIssued {

		return fmt.Errorf("Книга %s от автора %s %s и так уже в библиотеке.\n", b.Title, b.Author.Name, b.Author.Surname)
	}
	b.IsIssued = false
	b.ReaderID = 0
	fmt.Printf("Книга %s от автора %s %s возвращена в библиотеку.\n\n", b.Title, b.Author.Name, b.Author.Surname)
	return nil
}

//Методы читателей

func (r *Reader) String() string {
	return fmt.Sprintf("%d: %s. Активен: %v\n", r.ID, r.Name, r.IsActive)
}
