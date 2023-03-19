package domain

import (
	modelDomain "bookstore/bookstore-author-ms/internal/author/domain/model"
	"github.com/rs/zerolog/log"
)

type Repository interface {
	Create(author modelDomain.Author) (err error)
	Get(author modelDomain.Author) (result modelDomain.Author, err error)
	Update(author modelDomain.Author) (err error)
	Delete(author modelDomain.Author) (err error)
}

type AuthorService struct {
	repository Repository
}

func NewAuthorService(repository Repository) *AuthorService {
	return &AuthorService{repository}
}

func (as *AuthorService) Create(author modelDomain.Author) (err error) {
	err = as.repository.Create(author)
	if err != nil {
		return
	}

	log.Info().Str("id", author.Id).Msg("author has been created")

	return
}

func (as *AuthorService) Get(author modelDomain.Author) (result modelDomain.Author, err error) {
	result, err = as.repository.Get(author)
	if err != nil {
		return
	}

	log.Info().Str("id", result.Id).Msg("author has been obtained")

	return
}

func (as *AuthorService) Update(author modelDomain.Author) (err error) {
	err = as.repository.Update(author)
	if err != nil {
		return
	}

	log.Info().Str("id", author.Id).Msg("author has been updated")

	return
}

func (as *AuthorService) Delete(author modelDomain.Author) (err error) {
	err = as.repository.Delete(author)
	if err != nil {
		return
	}

	log.Info().Str("id", author.Id).Msg("author has been deleted")

	return
}
