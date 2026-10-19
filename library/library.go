package library

import (
	"fmt"
	"os"
	"strconv"
	"encoding/csv"
	"Lib/domain"
)

type Library struct {
	Readers []domain.Reader
	Books   []domain.Book

	lAddBookID   int
	lAddReadID int
}

// Иницализация

func CreateLibrary(books []domain.Book, readers []domain.Reader) *Library {
	library := Library{
		Readers:           readers,
		Books:             books,
		lAddBookID:   len(books) - 1,
		lAddReadID: len(readers) - 1,
	}
	fmt.Println("Новая библиотека была успешно создана!")
	return &library
}

// Обновление LastID У загруженной библиотеки

func (l *Library) UpdateLastID() {
	maxBookID := 0
	maxReadID := 0
	
	for _, book := range l.Books {
		if book.ID > maxBookID {
			maxBookID = book.ID
		}
	}

	for _, reader := range l.Readers {
		if reader.ID > maxReadID {
			maxReadID = reader.ID
		}
	}
	l.SetLAddBookID(maxBookID) 
	l.SetLAddReadID(maxReadID) 
}

// Методы библиотеки
// Методы вывода и ввода

func (l Library) GetLAddBookID() int {
	return l.lAddBookID
}

func (l *Library) SetLAddBookID( num int) {
	l.lAddBookID = num
}

func (l *Library) GetLAddReadID() int {
	return l.lAddReadID
}

func (l *Library) SetLAddReadID( num int) {
	l.lAddReadID = num
}

func (l *Library) GetAllBooks() {
	fmt.Print("Все книги:\n")
	for _, value := range l.Books {
		fmt.Println(value)
	}
}

func (l *Library) GetAllReaders() {
	fmt.Print("Все чители:")
	for _, value := range l.Readers {
		fmt.Println(value)
	}
}

// Методы поиска
// Поиск книги по ID
func (l *Library) GetBookById(BookId int) (*domain.Book, error) {
	for index := range l.Books {
		if BookId == l.Books[index].ID {
			fmt.Println("Книга найдена: ", l.Books[index].String())
			return &l.Books[index], nil
		}
	}
	return nil, fmt.Errorf("Книга не найдена. Просмотрите список всех книг и найдите нужный Вам ID: GetAllBooks()\n")
}

// Поиск книги по названию
func (l *Library) GetBookByTitle(BookTitle string) (*domain.Book, error) {
	for index := range l.Books {
		if l.Books[index].Title == BookTitle {
			fmt.Println("Книга найдена: \n", l.Books[index].String())
			return &l.Books[index], nil
		}
	}
	return nil, fmt.Errorf("\nКнига не найдена. Просмотрите список всех книг и найдите нужное Вам название: GetAllBooks()\n")
}

// Поиск читателя по ID
func (l *Library) GetReaderById(ReaderID int) (*domain.Reader, error) {
	for index := range l.Readers {
		if l.Readers[index].ID == ReaderID {
			fmt.Println("Читатель найден: \n", l.Readers[index].String())
			return &l.Readers[index], nil
		}
	}
	return nil, fmt.Errorf("Читатель не найден. Просмотрите список всех читателей и найдите нужный Вам ID: GetAllReaders()\n")
}

// Поиск читателя по имени
func (l *Library) GetReaderByName(ReaderName string) (*domain.Reader, error) {
	for index := range l.Readers {
		if l.Readers[index].Name == ReaderName {
			fmt.Println("Читатель найден: \n", l.Readers[index].String())
			return &l.Readers[index], nil
		}
	}
	return nil, fmt.Errorf("Читатель не найден. Просмотрите список всех читателей и найдите нужное Вам Имя: GetAllReaders()\n")
}

// Методы добавления

func (l *Library) AddBook(title string, author_name string, author_surname string) (*domain.Book, error) {
	if _, e := l.GetBookByTitle(title); e != nil {
		return nil, fmt.Errorf("Книга с таким названием уже есть в библиотеке.")
	}
	book := domain.Book{
		ID:       l.lAddBookID + 1,
		Title:    title,
		Author:   domain.Author{Name: author_name, Surname: author_surname},
		IsIssued: false,
		ReaderID: 0,
	}
	l.Books = append(l.Books, book)
	l.lAddBookID++
	fmt.Printf("Добавленная книга: %v\n", book.String())
	return &book, nil

}

// Еще методы
// Выдача книги
func (l *Library) IssueBookToReader(BookID int, ReaderID int) error {
	var book *domain.Book
	var reader *domain.Reader
	var err error
	fmt.Print("Выдача книги: \n")
	if book, err = l.GetBookById(BookID); err != nil {
		return fmt.Errorf("Некорректный ID книги")
	}
	if reader, err = l.GetReaderById(ReaderID); err != nil {
		return fmt.Errorf("Некорректный ID читателя")
	}
	if eror := book.IssueBook(*reader); eror != nil {
		return eror
	}

	fmt.Printf("Книга %s by %s %s выдана %s\n", book.Title, book.Author.Name, book.Author.Surname, reader.Name)
	return nil
}

func (l *Library) ReturnBook(BookID int) error {
	b, err := l.GetBookById(BookID)

	if err != nil {
		return err
	}

	err = b.ReturnBook()

	if err != nil {
		return err
	}
	return nil
}

// Методы работы с файлами

func (l *Library) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	writer.Write([]string{"id", "title", "authorName", "authorSurname", "isIssued", "readerID"})

	for _, book := range l.Books {
		id := strconv.Itoa(book.ID)
		isIssued := strconv.FormatBool(book.IsIssued)
		readerID := strconv.Itoa(book.ReaderID)
		writer.Write([]string{id, book.Title, book.Author.Name, book.Author.Surname, isIssued, readerID})
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}

func (l *Library) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	reader.Read()

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records {

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return err
		}

		isIssued, err := strconv.ParseBool(record[4])
		if err != nil {
			return err
		}

		readerID, err := strconv.Atoi(record[5])
		if err != nil {
			return err
		}

		l.Books = append(l.Books, domain.Book{
			ID:    id,
			Title: record[1],
			Author: domain.Author{
				Name:    record[2],
				Surname: record[3],
			},
			IsIssued: isIssued,
			ReaderID: readerID,
		})
	}

	return nil
}
