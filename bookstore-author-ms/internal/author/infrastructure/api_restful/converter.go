package api_restful

import (
	modelDomain "bookstore/bookstore-author-ms/internal/author/domain/model"
	"bookstore/bookstore-author-ms/internal/author/infrastructure/api_restful/openapi"
	"bookstore/internal/commons"
)

type APIConverter interface {
	APIToDomain(authorAPI openapi.AuthorRequestDTO) (modelDomain.Author, error)
	DomainsToAPI(authorList []modelDomain.Author) []openapi.AuthorResponseDTO
	DomainToAPI(authorDomain modelDomain.Author) openapi.AuthorResponseDTO
}

type AuthorAPIConverter struct {
}

func NewAuthorApiConverter() APIConverter {
	return &AuthorAPIConverter{}
}

func (ac *AuthorAPIConverter) APIToDomain(authorAPI openapi.AuthorRequestDTO) (modelDomain.Author, error) {
	bookID := commons.PointerToValue(authorAPI.BookId)

	return modelDomain.NewAuthor(authorAPI.Id, bookID, authorAPI.Name)
}

func (ac *AuthorAPIConverter) DomainsToAPI(authorsDomain []modelDomain.Author) []openapi.AuthorResponseDTO {
	resultList := make([]openapi.AuthorResponseDTO, len(authorsDomain))
	for i, author := range authorsDomain {
		resultList[i] = ac.DomainToAPI(author)
	}

	return resultList
}

func (ac *AuthorAPIConverter) DomainToAPI(authorDomain modelDomain.Author) openapi.AuthorResponseDTO {
	return openapi.AuthorResponseDTO{
		Id:     authorDomain.Id,
		BookId: &authorDomain.BookId,
		Name:   authorDomain.Name,
	}
}
