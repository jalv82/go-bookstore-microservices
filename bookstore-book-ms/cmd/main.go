package main

import (
	"bookstore/bookstore-book-ms/internal/book/domain"
	"bookstore/bookstore-book-ms/internal/book/domain/model"
	"bookstore/bookstore-book-ms/internal/book/infrastructure"
	database "bookstore/bookstore-book-ms/internal/book/infrastructure/database"
	"bookstore/internal/commons"
	"github.com/google/uuid"
)

func main() {
	databaseConfig := commons.ReadDatabaseConfig("bookstore-book-ms/config.yaml")
	bookSQLClient := database.NewBookSqlClient(databaseConfig)
	bookSQLConverter := database.NewBookSQLConverter()
	repository := infrastructure.NewRepository(*bookSQLClient, *bookSQLConverter)
	service := domain.NewBookService(repository)

	book := model.Book{
		Id:       uuid.NewString(),
		AuthorId: uuid.NewString(),
		Title:    "Go in Action",
	}

	err := service.Create(book)
	if err != nil {
		return
	}

	_, err = service.Get(book)
	if err != nil {
		return
	}

	book.Title = "Go in Action 2nd Edition"
	err = service.Update(book)
	if err != nil {
		return
	}

	err = service.Delete(book)
	if err != nil {
		return
	}
}
