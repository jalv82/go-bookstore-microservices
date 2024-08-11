package api_restful

import (
	"github.com/stretchr/testify/mock"

	modelDomain "bookstore/bookstore-author-ms/internal/author/domain/model"
	"bookstore/bookstore-author-ms/internal/author/infrastructure/api_restful/openapi"
)

type MockAuthorAPIConverter struct {
	mock.Mock
}

func (m *MockAuthorAPIConverter) APIToDomain(authorAPI openapi.AuthorRequestDTO) (modelDomain.Author, error) {
	args := m.Called(authorAPI)
	return args.Get(0).(modelDomain.Author), args.Error(1)
}

func (m *MockAuthorAPIConverter) DomainsToAPI(authorsDomain []modelDomain.Author) []openapi.AuthorResponseDTO {
	args := m.Called(authorsDomain)
	return args.Get(0).([]openapi.AuthorResponseDTO)
}

func (m *MockAuthorAPIConverter) DomainToAPI(authorDomain modelDomain.Author) openapi.AuthorResponseDTO {
	args := m.Called(authorDomain)
	return args.Get(0).(openapi.AuthorResponseDTO)
}
