package main

import (
	"bookstore/bookstore-author-ms/internal/author/application/crud"
	apiRestful "bookstore/bookstore-author-ms/internal/author/infrastructure/api_restful"
	"bookstore/bookstore-author-ms/internal/author/infrastructure/api_restful/openapi"
	httpServer "bookstore/bookstore-author-ms/internal/author/infrastructure/http_server"
	sqlDatabase "bookstore/bookstore-author-ms/internal/author/infrastructure/sql_database"
	"bookstore/internal/commons"
)

func main() {
	databaseConfig := commons.ReadDatabaseConfig("bookstore-author-ms/config.yaml")
	postgreSQLClient := commons.NewPostgreSQLClient(databaseConfig)

	authorSQLClient := sqlDatabase.NewAuthorSQLClient(postgreSQLClient)
	authorSQLConverter := sqlDatabase.NewAuthorSQLConverter()
	repository := sqlDatabase.NewRepository(*authorSQLClient, authorSQLConverter)

	service := crud.NewAuthorService(repository)
	authorAPIConverter := apiRestful.NewAuthorApiConverter()
	controller := apiRestful.NewAuthorHttpController(&service, authorAPIConverter)

	httpServerConfig := httpServer.ReadHttpServerConfig("bookstore-author-ms/config.yaml")
	httpServerInstance := httpServer.NewHttpServer(httpServerConfig)
	openapi.RegisterHandlers(httpServerInstance.Server, &controller)

	err := httpServerInstance.Up()
	if err != nil {
		return
	}
}
