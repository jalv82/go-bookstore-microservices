package domain

import (
	modelDomain "bookstore/bookstore-book-ms/internal/book/domain/model"
	"github.com/rs/zerolog/log"
)

type Repository interface {
	Create(book modelDomain.Book) (err error)
	Get(book modelDomain.Book) (result modelDomain.Book, err error)
	Update(book modelDomain.Book) (err error)
	Delete(book modelDomain.Book) (err error)
}

type BookService struct {
	repository Repository
}

func NewBookService(repository Repository) *BookService {
	return &BookService{repository}
}

func (as *BookService) Create(book modelDomain.Book) (err error) {
	err = as.repository.Create(book)
	if err != nil {
		return
	}

	log.Info().Str("id", book.Id).Msg("book has been created")

	return
}

func (as *BookService) Get(book modelDomain.Book) (result modelDomain.Book, err error) {
	result, err = as.repository.Get(book)
	if err != nil {
		return
	}

	log.Info().Str("id", result.Id).Msg("book has been obtained")

	return
}

func (as *BookService) Update(book modelDomain.Book) (err error) {
	err = as.repository.Update(book)
	if err != nil {
		return
	}

	log.Info().Str("id", book.Id).Msg("book has been updated")

	return
}

func (as *BookService) Delete(book modelDomain.Book) (err error) {
	err = as.repository.Delete(book)
	if err != nil {
		return
	}

	log.Info().Str("id", book.Id).Msg("book has been deleted")

	return
}
