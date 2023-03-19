package infrastructure

import (
	"errors"

	modelDomain "bookstore/bookstore-book-ms/internal/book/domain/model"
	"bookstore/bookstore-book-ms/internal/book/infrastructure/database"
	"github.com/rs/zerolog/log"
)

var (
	ErrCreateBook = errors.New("error creating book")
	ErrGetBook    = errors.New("error getting book")
	ErrUpdateBook = errors.New("error updating book")
	ErrDeleteBook = errors.New("error deleting book")
)

type BookRepository struct {
	sqlClient database.BookSQLClient
	converter database.BookSQLConverter
}

func NewRepository(bookSQLClient database.BookSQLClient, bookSQLConverter database.BookSQLConverter) *BookRepository {
	return &BookRepository{bookSQLClient, bookSQLConverter}
}
func (ar *BookRepository) Create(book modelDomain.Book) (err error) {
	bookDB := ar.converter.DomainToDB(book)

	err = ar.sqlClient.Create(&bookDB)
	if err != nil {
		err = ErrCreateBook
		log.Error().Str("id", book.Id).Msg(err.Error())
	}

	return
}

func (ar *BookRepository) Get(book modelDomain.Book) (result modelDomain.Book, err error) {
	bookDB := ar.converter.DomainToDB(book)

	resultDB, err := ar.sqlClient.Get(&bookDB)
	if err != nil {
		err = ErrGetBook
		log.Error().Str("id", book.Id).Msg(err.Error())
	}

	result = ar.converter.DBtoDomain(*resultDB)
	return
}

func (ar *BookRepository) Update(book modelDomain.Book) (err error) {
	bookDB := ar.converter.DomainToDB(book)

	err = ar.sqlClient.Update(&bookDB)
	if err != nil {
		err = ErrUpdateBook
		log.Error().Str("id", book.Id).Msg(err.Error())
	}

	return
}

func (ar *BookRepository) Delete(book modelDomain.Book) (err error) {
	bookDB := ar.converter.DomainToDB(book)

	err = ar.sqlClient.Delete(&bookDB)
	if err != nil {
		err = ErrDeleteBook
		log.Error().Str("id", book.Id).Msg(err.Error())
	}

	return
}
