package main

import (
	"Lib/cmd/api"
	"Lib/domain"
	"Lib/library"
	"Lib/storage"
	"flag"
	"fmt"
	"log"
)

func showFunctional(lib *library.Library) {
	// Вывод всех книг до добавления
	lib.GetAllBooks()
	// Добавление новой книги
	addedBook, err := lib.AddBook("Преступление и наказание", "Фёдор", "Достоевский")
	if err != nil {
		fmt.Println("Ошибка добавления книги:", err)
	} else {
		fmt.Println("Книга успешно добавлена:", addedBook.Title)
	}
	// Вывод всех книг после добавления
	lib.GetAllBooks()
	// Поиск существующей книги по ID
	foundBookByID, err := lib.GetBookById(2)
	if err != nil {
		fmt.Println("Ошибка поиска книги по ID:", err)
	} else {
		fmt.Println("Найдена книга:", foundBookByID.Title)
	}
	// Поиск несуществующей книги по ID
	missingBookByID, err := lib.GetBookById(999)
	if err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	} else {
		fmt.Println("Найдена книга:", missingBookByID.Title)
	}
	// Поиск существующей книги по названию
	foundBookByTitle, err := lib.GetBookByTitle("Таинственный остров")
	if err != nil {
		fmt.Println("Ошибка поиска книги по названию:", err)
	} else {
		fmt.Println("Найдена книга:", foundBookByTitle.Title)
	}
	// Поиск несуществующей книги по названию
	missingBookByTitle, err := lib.GetBookByTitle("Гарри Поттер")
	if err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	} else {
		fmt.Println("Найдена книга:", missingBookByTitle.Title)
	}
	// Поиск существующего читателя по ID
	foundReaderByID, err := lib.GetReaderById(3)
	if err != nil {
		fmt.Println("Ошибка поиска читателя по ID:", err)
	} else {
		fmt.Println("Найден читатель:", foundReaderByID.Name)
	}
	// Поиск несуществующего читателя по ID
	missingReaderByID, err := lib.GetReaderById(999)
	if err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	} else {
		fmt.Println("Найден читатель:", missingReaderByID.Name)
	}
	// Поиск существующего читателя по имени
	foundReaderByName, err := lib.GetReaderByName("Милана")
	if err != nil {
		fmt.Println("Ошибка поиска читателя по имени:", err)
	} else {
		fmt.Println("Найден читатель:", foundReaderByName.Name)
	}
	// Поиск несуществующего читателя по имени
	missingReaderByName, err := lib.GetReaderByName("Константин")
	if err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	} else {
		fmt.Println("Найден читатель:", missingReaderByName.Name)
	}
	// Попытка выдать книгу неактивному читателю
	err = lib.IssueBookToReader(2, 3)
	if err != nil {
		fmt.Println("Ожидаемая ошибка выдачи:", err)
	}
	// Активация читателя
	reader, err := lib.GetReaderById(3)
	if err != nil {
		fmt.Println("Ошибка поиска читателя:", err)
		return
	}
	reader.IsActive = true
	fmt.Println("Читатель", reader.Name, "активирован")
	// Повторная попытка выдачи книги
	err = lib.IssueBookToReader(2, 3)
	if err != nil {
		fmt.Println("Ошибка выдачи книги:", err)
	} else {
		fmt.Println("Книга успешно выдана читателю")
	}
	// Попытка повторно выдать уже выданную книгу
	err = lib.IssueBookToReader(2, 4)
	if err != nil {
		fmt.Println("Ожидаемая ошибка повторной выдачи:", err)
	}
	// Возврат выданной книги
	err = lib.ReturnBook(2)
	if err != nil {
		fmt.Println("Ошибка возврата книги:", err)
	} else {
		fmt.Println("Книга успешно возвращена в библиотеку")
	}
	// Попытка повторного возврата
	err = lib.ReturnBook(2)
	if err != nil {
		fmt.Println("Ожидаемая ошибка повторного возврата:", err)
	}
}

func main() {

	// Регистрация флагов командной строки
	saveFlag := flag.Bool("save", false, "Сохранить библиотеку в JSON")
	loadFlag := flag.Bool("load", false, "Загрузить библиотеку из JSON")
	showBasicsFlag := flag.Bool("showBasics", false, "Показать базовый функционал")
	flag.Parse()
	// Создание начальных читателей
	readers := []domain.Reader{
		{ID: 1, Name: "Мигель", IsActive: false},
		{ID: 2, Name: "Александр", IsActive: false},
		{ID: 3, Name: "Анатолий", IsActive: false},
		{ID: 4, Name: "Милана", IsActive: false},
		{ID: 5, Name: "Игорь", IsActive: false},
		{ID: 6, Name: "Алана", IsActive: false},
		{ID: 7, Name: "Виктория", IsActive: false},
	}
	// Создание начальных книг
	books := []domain.Book{
		{
			ID:     1,
			Title:  "Вокруг света за 80 дней",
			Author: domain.Author{Name: "Жуль", Surname: "Верн"},
		},
		{
			ID:     2,
			Title:  "Таинственный остров",
			Author: domain.Author{Name: "Жуль", Surname: "Верн"},
		},
		{
			ID:     3,
			Title:  "Остров сокровищ",
			Author: domain.Author{Name: "Роберт", Surname: "Стивенсон"},
		},
		{
			ID:     4,
			Title:  "2:36 по Аляске",
			Author: domain.Author{Name: "Анастасия", Surname: "Гор"},
		},
		{
			ID:     5,
			Title:  "Менталитет Мамбы",
			Author: domain.Author{Name: "Коби", Surname: "Брайант"},
		},
	}
	// Создание JSON-хранилища через интерфейс
	var dataStorage storage.Storable = &storage.JSONStorage{}
	// Инициализация библиотеки
	var lib *library.Library
	var err error
	if *loadFlag {
		// Загружаем библиотеку из JSON
		lib, err = dataStorage.Load("books.json")
		if err != nil {
			log.Fatal("Ошибка загрузки библиотеки:", err)
		}
		fmt.Println("Библиотека успешно загружена из JSON")
	} else {
		// Создаём библиотеку с начальными данными
		lib = library.CreateLibrary(books, readers)
		fmt.Println("Библиотека создана с начальными данными")
	}
	// Информация о библиотеке
	fmt.Println("Книг в библиотеке:", len(lib.Books))
	fmt.Println("Читателей в библиотеке:", len(lib.Readers))
	// Демонстрация базового функционала
	if *showBasicsFlag {
		showFunctional(lib)
	}
	// Сохранение библиотеки
	if *saveFlag {
		if err := dataStorage.Save("books.json", lib); err != nil {
			log.Fatal("Ошибка сохранения библиотеки:", err)
		}
		fmt.Println("Библиотека успешно сохранена в books.json")
	}
	// Запуск HTTP-сервера
	fmt.Println("HTTP-сервер запускается на порту 8080")
	api.Init()
}
