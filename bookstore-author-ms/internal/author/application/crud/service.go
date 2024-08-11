package crud

import (
	"github.com/rs/zerolog/log"

	modelDomain "bookstore/bookstore-author-ms/internal/author/domain"
	"bookstore/bookstore-author-ms/internal/author/domain/model"
)

type AuthorService struct {
	repository modelDomain.CRUDPort
}

var _ modelDomain.CRUDPort = (*AuthorService)(nil)

func NewAuthorService(repository modelDomain.CRUDPort) AuthorService {
	return AuthorService{repository}
}

func (as *AuthorService) Create(author model.Author) error {
	err := as.repository.Create(author)
	if err != nil {
		return err
	}

	log.Info().Str("id", author.Id).Msg("author has been created")

	return err
}

func (as *AuthorService) Get(author model.Author) (model.Author, error) {
	result, err := as.repository.Get(author)
	if err != nil {
		return model.Author{}, err
	}

	log.Info().Str("id", result.Id).Msg("author has been obtained")

	return result, nil
}

func (as *AuthorService) GetAll() []model.Author {
	result := as.repository.GetAll()

	log.Info().Msg("authors has been obtained")

	return result
}

func (as *AuthorService) Update(author model.Author) error {
	err := as.repository.Update(author)
	if err != nil {
		return err
	}

	log.Info().Str("id", author.Id).Msg("author has been updated")

	return nil
}

func (as *AuthorService) Delete(author model.Author) error {
	err := as.repository.Delete(author)
	if err != nil {
		return err
	}

	log.Info().Str("id", author.Id).Msg("author has been deleted")

	return nil
}
