package main

import (
	"bookstore/bookstore-author-ms/internal/author/domain"
	"bookstore/bookstore-author-ms/internal/author/domain/model"
	"bookstore/bookstore-author-ms/internal/author/infrastructure"
	database "bookstore/bookstore-author-ms/internal/author/infrastructure/database"
	"bookstore/internal/commons"
	"github.com/google/uuid"
)

func main() {
	readDatabaseConfig := commons.ReadDatabaseConfig("bookstore-author-ms/config.yaml")
	authorSQLClient := database.NewAuthorSqlClient(readDatabaseConfig)
	authorSQLConverter := database.NewAuthorSQLConverter()
	repository := infrastructure.NewRepository(*authorSQLClient, *authorSQLConverter)
	service := domain.NewAuthorService(repository)

	author := model.Author{
		Id:     uuid.NewString(),
		BookId: uuid.NewString(),
		Name:   "William",
	}

	err := service.Create(author)
	if err != nil {
		return
	}

	_, err = service.Get(author)
	if err != nil {
		return
	}

	author.Name = "William Kennedy"
	err = service.Update(author)
	if err != nil {
		return
	}

	err = service.Delete(author)
	if err != nil {
		return
	}
}
