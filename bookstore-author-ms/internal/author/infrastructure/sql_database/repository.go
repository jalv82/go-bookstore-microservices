package sql_database

import (
	"github.com/rs/zerolog/log"

	"bookstore/bookstore-author-ms/internal/author/domain"
	modelDomain "bookstore/bookstore-author-ms/internal/author/domain/model"
)

type AuthorRepository struct {
	sqlClient AuthorSQLClient
	converter AuthorSQLConverter
}

var _ domain.CRUDPort = (*AuthorRepository)(nil)

func NewRepository(authorSQLClient AuthorSQLClient, authorSQLConverter AuthorSQLConverter) *AuthorRepository {
	return &AuthorRepository{authorSQLClient, authorSQLConverter}
}

func (ar *AuthorRepository) Create(author modelDomain.Author) error {
	authorDB := ar.converter.DomainToDB(author)

	err := ar.sqlClient.Create(&authorDB)
	if err != nil {
		log.Error().Str("id", author.Id).Msg(err.Error())

		return err
	}

	return nil
}

func (ar *AuthorRepository) Get(author modelDomain.Author) (modelDomain.Author, error) {
	authorDB := ar.converter.DomainToDB(author)

	err := ar.sqlClient.Get(&authorDB)
	if err != nil {
		log.Error().Str("id", author.Id).Msg(err.Error())

		return modelDomain.Author{}, err
	}

	return ar.converter.DBToDomain(authorDB), nil
}

func (ar *AuthorRepository) GetAll() []modelDomain.Author {
	authorsDB := ar.sqlClient.GetAll()

	return ar.converter.DBsToDomain(authorsDB)
}

func (ar *AuthorRepository) Update(author modelDomain.Author) error {
	authorDB := ar.converter.DomainToDB(author)

	err := ar.sqlClient.Update(&authorDB)
	if err != nil {
		log.Error().Str("id", author.Id).Msg(err.Error())

		return err
	}

	return nil
}

func (ar *AuthorRepository) Delete(author modelDomain.Author) error {
	authorDB := ar.converter.DomainToDB(author)

	err := ar.sqlClient.Delete(&authorDB)
	if err != nil {
		log.Error().Str("id", author.Id).Msg(err.Error())

		return err
	}

	return nil
}
