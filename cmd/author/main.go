package main

import (
	commons "bookstore/internal/commons/postgres"
	"bookstore/internal/core/author/domain"
	modelDomain "bookstore/internal/core/author/domain/model"
	infra "bookstore/internal/core/author/infrastructure"
	"bookstore/internal/core/author/infrastructure/database"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	readDatabaseConfig := readDatabaseConfig()
	authorSQLClient := database.NewAuthorSqlClient(readDatabaseConfig)
	authorSQLConverter := database.NewAuthorSQLConverter()
	repository := infra.NewRepository(*authorSQLClient, *authorSQLConverter)
	service := domain.NewAuthorService(repository)

	author := modelDomain.Author{
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

	// err = service.Delete(author)
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
		Driver:   viper.GetString("database-author.driver"),
		User:     viper.GetString("database-author.user"),
		Password: viper.GetString("database-author.password"),
		Host:     viper.GetString("database-author.host"),
		Port:     viper.GetString("database-author.port"),
		Schema:   viper.GetString("database-author.schema"),
	}

	return
}
