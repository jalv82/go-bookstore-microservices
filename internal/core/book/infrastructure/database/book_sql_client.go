package database

import (
	commons "bookstore/internal/commons/postgres"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type BookSQLClient struct {
	db *gorm.DB
}

func NewBookSQLClient(databaseConfig *commons.DatabaseConfig) *BookSQLClient {
	db, err := commons.NewGormClient(databaseConfig)
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	return &BookSQLClient{db}
}

func (a *BookSQLClient) Create(book *Book) error {
	return a.db.Table(book.TableName()).Create(book).Error
}

func (a *BookSQLClient) Get(book *Book) (result *Book, err error) {
	err = a.db.Table(book.TableName()).First(&result).Error
	return
}

func (a *BookSQLClient) Update(book *Book) error {
	return a.db.Table(book.TableName()).Save(&book).Error
}

func (a *BookSQLClient) Delete(book *Book) error {
	return a.db.Table(book.TableName()).Delete(&book).Error
}
