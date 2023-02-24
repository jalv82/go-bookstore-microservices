package database

import (
	modelDomain "bookstore/internal/core/author/domain/model"
)

type AuthorSQLConverter struct {
}

func NewAuthorSQLConverter() *AuthorSQLConverter {
	return &AuthorSQLConverter{}
}

func (ac *AuthorSQLConverter) DomainToDB(authorDomain modelDomain.Author) (authorDB Author) {
	authorDB = Author{
		Id:     authorDomain.Id,
		BookId: authorDomain.BookId,
		Name:   authorDomain.Name,
	}

	return
}

func (ac *AuthorSQLConverter) DBtoDomain(authorDB Author) (authorDomain modelDomain.Author) {
	authorDomain = modelDomain.Author{
		Id:     authorDB.Id,
		BookId: authorDB.BookId,
		Name:   authorDB.Name,
	}

	return
}
