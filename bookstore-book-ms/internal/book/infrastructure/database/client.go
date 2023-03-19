package database

import (
	"bookstore/internal/commons"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

type BookSQLClient struct {
	db *gorm.DB
}

func NewBookSqlClient(databaseConfig *commons.DatabaseConfig) *BookSQLClient {
	db := commons.NewSqlClient(databaseConfig)

	return &BookSQLClient{db}
}

func NewMockBookSqlClient() (sqlmock.Sqlmock, *BookSQLClient) {
	db, mock := commons.NewMockSqlClient()

	return mock, &BookSQLClient{db}
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
