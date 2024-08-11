package sql_database

import (
	"errors"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"

	"bookstore/bookstore-author-ms/internal/author/infrastructure/sql_database/model"
)

const (
	errorRecordExists   = "record already exists with id: %s"
	errorRecordNotFound = "record not found with id: %s"
)

type AuthorSQLClient struct {
	db *gorm.DB
}

func NewAuthorSQLClient(db *gorm.DB) *AuthorSQLClient {
	return &AuthorSQLClient{db}
}

func NewMockAuthorSQLClient(db *gorm.DB, mock sqlmock.Sqlmock) (sqlmock.Sqlmock, *AuthorSQLClient) {
	return mock, &AuthorSQLClient{db}
}

func (a *AuthorSQLClient) Create(author *model.Author) error {
	err := a.db.Table(model.Authors).Create(author).Error
	if err != nil {
		return fmt.Errorf(errorRecordExists, author.Id)
	}

	return nil
}

func (a *AuthorSQLClient) Get(author *model.Author) error {
	return a.db.Table(model.Authors).First(&author).Error
}

func (a *AuthorSQLClient) GetAll() []model.Author {
	var authors []model.Author
	a.db.Table(model.Authors).Find(&authors)

	return authors
}

// Update updates an author in the database based on the provided author object.
func (a *AuthorSQLClient) Update(author *model.Author) error {
	authorToFind := *author
	err := a.db.Table(model.Authors).First(&authorToFind).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf(errorRecordNotFound, authorToFind.Id)
		}

		return err
	}

	return a.db.Table(model.Authors).Save(&author).Error
}

func (a *AuthorSQLClient) Delete(author *model.Author) error {
	db := a.db.Table(model.Authors).Delete(&author)
	if db.RowsAffected == 0 {
		return fmt.Errorf(errorRecordNotFound, author.Id)
	}

	return nil
}
