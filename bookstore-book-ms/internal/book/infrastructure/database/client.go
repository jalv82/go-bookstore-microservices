package database

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"

	"bookstore/internal/commons"
)

type BookSQLClient struct {
	db *gorm.DB
}

func NewBookSqlClient(databaseConfig *commons.DatabaseConfig) *BookSQLClient {
	db := commons.NewPostgreSQLClient(databaseConfig)

	return &BookSQLClient{db}
}

func NewMockBookSqlClient() (sqlmock.Sqlmock, *BookSQLClient) {
	db, mock := commons.NewPostgreSQLClientMock()

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
