package main

import (
	commons "bookstore/internal/commons/postgres"
	"bookstore/internal/core/book/domain"
	modelDomain "bookstore/internal/core/book/domain/model"
	infra "bookstore/internal/core/book/infrastructure"
	"bookstore/internal/core/book/infrastructure/database"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	databaseConfig := readDatabaseConfig()
	bookSQLClient := database.NewBookSQLClient(databaseConfig)
	bookSQLConverter := database.NewBookSQLConverter()
	repository := infra.NewRepository(*bookSQLClient, *bookSQLConverter)
	service := domain.NewBookService(repository)

	book := modelDomain.Book{
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

	// err = service.Delete(book)
	// if err != nil {
	// 	return
	// }
}

func readDatabaseConfig() (databaseConfig *commons.DatabaseConfig) {
	viper.SetConfigFile("config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	databaseConfig = &commons.DatabaseConfig{
		Driver:   viper.GetString("database-book.driver"),
		User:     viper.GetString("database-book.user"),
		Password: viper.GetString("database-book.password"),
		Host:     viper.GetString("database-book.host"),
		Port:     viper.GetString("database-book.port"),
		Schema:   viper.GetString("database-book.schema"),
	}

	return
}
