package api_restful_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	modelDomain "bookstore/bookstore-author-ms/internal/author/domain/model"
	"bookstore/bookstore-author-ms/internal/author/infrastructure/api_restful"
	"bookstore/bookstore-author-ms/internal/author/infrastructure/api_restful/openapi"
	"bookstore/internal/commons"
)

func TestAuthorAPIConverter_APIToDomain(t *testing.T) {
	t.Parallel()

	// Given
	authorAPI := openapi.AuthorRequestDTO{
		Id:     uuid.NewString(),
		BookId: commons.ValueOfPointer(uuid.NewString()),
		Name:   "William Kennedy",
	}

	// When
	converter := api_restful.NewAuthorApiConverter()
	authorDomain, err := converter.APIToDomain(authorAPI)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, authorAPI.Id, authorDomain.Id)
	assert.Equal(t, authorAPI.Name, authorDomain.Name)
	assert.Equal(t, *authorAPI.BookId, authorDomain.BookId)
}

func TestAuthorAPIConverter_DomainsToAPI(t *testing.T) {
	t.Parallel()

	// Given
	authorsDomain := []modelDomain.Author{
		{
			Id:     uuid.NewString(),
			BookId: uuid.NewString(),
			Name:   "William Kennedy",
		},
	}

	// When
	converter := api_restful.NewAuthorApiConverter()
	authorsAPI := converter.DomainsToAPI(authorsDomain)

	// Then
	assert.Equal(t, authorsDomain[0].Id, authorsAPI[0].Id)
	assert.Equal(t, authorsDomain[0].Name, authorsAPI[0].Name)
	assert.Equal(t, authorsDomain[0].BookId, *authorsAPI[0].BookId)
}

func TestAuthorAPIConverter_DomainToAPI(t *testing.T) {
	t.Parallel()

	// Given
	authorDomain := modelDomain.Author{
		Id:     uuid.NewString(),
		BookId: uuid.NewString(),
		Name:   "William Kennedy",
	}

	// When
	converter := api_restful.NewAuthorApiConverter()
	authorAPI := converter.DomainToAPI(authorDomain)

	// Then
	assert.Equal(t, authorDomain.Id, authorAPI.Id)
	assert.Equal(t, authorDomain.Name, authorAPI.Name)
	assert.Equal(t, authorDomain.BookId, *authorAPI.BookId)
}
