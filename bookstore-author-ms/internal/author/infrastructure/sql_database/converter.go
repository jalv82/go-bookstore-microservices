package sql_database

import (
	modelDomain "bookstore/bookstore-author-ms/internal/author/domain/model"
	modelDB "bookstore/bookstore-author-ms/internal/author/infrastructure/sql_database/model"
)

type SQLConverter interface {
	DomainToDB(authorDomain modelDomain.Author) modelDB.Author
	DBsToDomain(authorsDB []modelDB.Author) []modelDomain.Author
	DBToDomain(authorDB modelDB.Author) modelDomain.Author
}

type AuthorSQLConverter struct {
}

func NewAuthorSQLConverter() AuthorSQLConverter {
	return AuthorSQLConverter{}
}

func (ac *AuthorSQLConverter) DomainToDB(authorDomain modelDomain.Author) modelDB.Author {
	return modelDB.Author{
		Id:     authorDomain.Id,
		BookId: authorDomain.BookId,
		Name:   authorDomain.Name,
	}
}

func (ac *AuthorSQLConverter) DBsToDomain(authorsDB []modelDB.Author) []modelDomain.Author {
	authorList := make([]modelDomain.Author, len(authorsDB))

	for i, author := range authorsDB {
		authorList[i] = ac.DBToDomain(author)
	}

	return authorList
}

func (ac *AuthorSQLConverter) DBToDomain(authorDB modelDB.Author) modelDomain.Author {
	return modelDomain.Author{
		Id:     authorDB.Id,
		BookId: authorDB.BookId,
		Name:   authorDB.Name,
	}
}
