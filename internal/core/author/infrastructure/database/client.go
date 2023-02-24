package database

import (
	commons "bookstore/internal/commons/postgres"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type AuthorSQLClient struct {
	db *gorm.DB
}

func NewAuthorSqlClient(databaseConfig *commons.DatabaseConfig) *AuthorSQLClient {
	db, err := commons.NewSqlClient(databaseConfig)
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	return &AuthorSQLClient{db}
}

func NewMockAuthorSqlClient() (sqlmock.Sqlmock, *AuthorSQLClient) {
	db, mock, err := commons.NewMockSqlClient()
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	return mock, &AuthorSQLClient{db}
}

func (a *AuthorSQLClient) Create(author *Author) error {
	return a.db.Table(author.TableName()).Create(author).Error
}

func (a *AuthorSQLClient) Get(author *Author) (result *Author, err error) {
	err = a.db.Table(author.TableName()).First(&result).Error
	return
}

func (a *AuthorSQLClient) Update(author *Author) error {
	return a.db.Table(author.TableName()).Save(&author).Error
}

func (a *AuthorSQLClient) Delete(author *Author) error {
	return a.db.Table(author.TableName()).Delete(&author).Error
}
