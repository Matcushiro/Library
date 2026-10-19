package storage

import (
    "encoding/json"
    "os"

    "Lib/domain"
    "Lib/library"
)

// Интерфейс хранилища
type Storable interface {
    Save(filename string, lib *library.Library) error
    Load(filename string) (*library.Library, error)
}

// Структура JSON-хранилища
type JSONStorage struct{}

// Вспомогательная структура для сериализации
type storageData struct {
    Books   []domain.Book   `json:"books"`
    Readers []domain.Reader `json:"readers"`
}

// Сохранение библиотеки
func (s *JSONStorage) Save(
    filePath string,
    lib *library.Library,
) error {

    data := storageData{
        Books:   lib.Books,
        Readers: lib.Readers,
    }

    dataJSON, err := json.MarshalIndent(data, "", "  ")

    if err != nil {
        return err
    }

    return os.WriteFile(filePath, dataJSON, 0644)
}

// Загрузка библиотеки
func (s *JSONStorage) Load(
    filePath string,
) (*library.Library, error) {

    dataJSON, err := os.ReadFile(filePath)

    if err != nil {
        return nil, err
    }

    var data storageData

    if err := json.Unmarshal(dataJSON, &data); err != nil {
        return nil, err
    }

    lib := library.CreateLibrary(data.Books, data.Readers)

    lib.UpdateLastID()

    return lib, nil
}