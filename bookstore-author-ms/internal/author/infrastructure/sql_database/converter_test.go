package sql_database_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	modelDomain "bookstore/bookstore-author-ms/internal/author/domain/model"
	"bookstore/bookstore-author-ms/internal/author/infrastructure/sql_database"
	modelDB "bookstore/bookstore-author-ms/internal/author/infrastructure/sql_database/model"
)

func TestAuthorAPIConverter_DomainToDB(t *testing.T) {
	t.Parallel()

	// Given
	authorDomain := modelDomain.Author{
		Id:     uuid.NewString(),
		BookId: uuid.NewString(),
		Name:   "William Kennedy",
	}

	// When
	converter := sql_database.NewAuthorSQLConverter()
	authorDB := converter.DomainToDB(authorDomain)

	// Then
	assert.Equal(t, authorDB.Id, authorDomain.Id)
	assert.Equal(t, authorDB.Name, authorDomain.Name)
	assert.Equal(t, authorDB.BookId, authorDomain.BookId)
}

func TestAuthorAPIConverter_DBsToDomain(t *testing.T) {
	t.Parallel()

	// Given
	authorsDB := []modelDB.Author{
		{
			Id:     uuid.NewString(),
			BookId: uuid.NewString(),
			Name:   "William Kennedy",
		},
	}

	// When
	converter := sql_database.NewAuthorSQLConverter()
	authorsDomain := converter.DBsToDomain(authorsDB)

	// Then
	assert.Equal(t, authorsDomain[0].Id, authorsDB[0].Id)
	assert.Equal(t, authorsDomain[0].Name, authorsDB[0].Name)
	assert.Equal(t, authorsDomain[0].BookId, authorsDB[0].BookId)
}

func TestAuthorAPIConverter_DBToDomain(t *testing.T) {
	t.Parallel()

	// Given
	authorDB := modelDB.Author{
		Id:     uuid.NewString(),
		BookId: uuid.NewString(),
		Name:   "William Kennedy",
	}

	// When
	converter := sql_database.NewAuthorSQLConverter()
	authorDomain := converter.DBToDomain(authorDB)

	// Then
	assert.Equal(t, authorDB.Id, authorDomain.Id)
	assert.Equal(t, authorDB.Name, authorDomain.Name)
	assert.Equal(t, authorDB.BookId, authorDomain.BookId)
}
