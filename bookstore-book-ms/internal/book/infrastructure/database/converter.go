package database

import (
	modelDomain "bookstore/bookstore-book-ms/internal/book/domain/model"
)

type BookSQLConverter struct {
}

func NewBookSQLConverter() *BookSQLConverter {
	return &BookSQLConverter{}
}

func (ac *BookSQLConverter) DomainToDB(bookDomain modelDomain.Book) (bookDB Book) {
	bookDB = Book{
		Id:       bookDomain.Id,
		AuthorId: bookDomain.AuthorId,
		Title:    bookDomain.Title,
	}

	return
}

func (ac *BookSQLConverter) DBtoDomain(bookDB Book) (bookDomain modelDomain.Book) {
	bookDomain = modelDomain.Book{
		Id:       bookDB.Id,
		AuthorId: bookDB.AuthorId,
		Title:    bookDB.Title,
	}

	return
}
