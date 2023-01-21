package infrastructure

import (
	"errors"

	modelDomain "bookstore/internal/core/author/domain/model"
	"bookstore/internal/core/author/infrastructure/database"
	"github.com/rs/zerolog/log"
)

var (
	ErrCreateAuthor = errors.New("error creating author")
	ErrGetAuthor    = errors.New("error getting author")
	ErrUpdateAuthor = errors.New("error updating author")
	ErrDeleteAuthor = errors.New("error deleting author")
)

type AuthorRepository struct {
	sqlClient database.AuthorSQLClient
	converter database.AuthorSQLConverter
}

func NewRepository(authorSQLClient database.AuthorSQLClient, authorSQLConverter database.AuthorSQLConverter) *AuthorRepository {
	return &AuthorRepository{authorSQLClient, authorSQLConverter}
}

func (ar *AuthorRepository) Create(author modelDomain.Author) (err error) {
	authorDB := ar.converter.DomainToDB(author)

	err = ar.sqlClient.Create(&authorDB)
	if err != nil {
		err = ErrCreateAuthor
		log.Error().Str("id", author.Id).Msg(err.Error())
	}

	return
}

func (ar *AuthorRepository) Get(author modelDomain.Author) (result modelDomain.Author, err error) {
	authorDB := ar.converter.DomainToDB(author)

	resultDB, err := ar.sqlClient.Get(&authorDB)
	if err != nil {
		err = ErrGetAuthor
		log.Error().Str("id", author.Id).Msg(err.Error())
	}

	result = ar.converter.DBtoDomain(*resultDB)
	return
}

func (ar *AuthorRepository) Update(author modelDomain.Author) (err error) {
	authorDB := ar.converter.DomainToDB(author)

	err = ar.sqlClient.Update(&authorDB)
	if err != nil {
		err = ErrUpdateAuthor
		log.Error().Str("id", author.Id).Msg(err.Error())
	}

	return
}

func (ar *AuthorRepository) Delete(author modelDomain.Author) (err error) {
	authorDB := ar.converter.DomainToDB(author)

	err = ar.sqlClient.Delete(&authorDB)
	if err != nil {
		err = ErrDeleteAuthor
		log.Error().Str("id", author.Id).Msg(err.Error())
	}

	return
}
