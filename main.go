package main

import "fmt"

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
