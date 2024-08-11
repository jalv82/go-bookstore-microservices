package api_restful

import (
	"errors"
	"net/http"

	"bookstore/bookstore-author-ms/internal/author/domain"
	"bookstore/bookstore-author-ms/internal/author/infrastructure/api_restful/openapi"
	server "bookstore/bookstore-author-ms/internal/author/infrastructure/http_server"
)

const indent = "  "

var (
	ErrCreateAuthor = errors.New("error creating Author")
	ErrGetAuthor    = errors.New("error getting Author")
	ErrUpdateAuthor = errors.New("error updating Author")
	ErrDeleteAuthor = errors.New("error deleting Author")
)

type AuthorController struct {
	service   domain.CRUDPort
	converter APIConverter
}

func NewAuthorHttpController(service domain.CRUDPort, converter APIConverter) AuthorController {
	return AuthorController{
		service:   service,
		converter: converter,
	}
}

func (ac *AuthorController) CreateAuthor(httpServerCtx server.HttpServerCtx) error {
	authorRequestDTO := &openapi.AuthorRequestDTO{}
	err := httpServerCtx.Bind(authorRequestDTO)
	if err != nil {
		return err
	}

	author, err := ac.converter.APIToDomain(*authorRequestDTO)
	if err != nil {
		return ac.customError(httpServerCtx, ErrCreateAuthor, err)
	}

	err = ac.service.Create(author)
	if err != nil {
		return ac.customError(httpServerCtx, ErrCreateAuthor, err)
	}

	return httpServerCtx.JSON(http.StatusCreated, ac.converter.DomainToAPI(author))
}

func (ac *AuthorController) GetAuthor(httpServerCtx server.HttpServerCtx, authorId string) error {
	authorRequestDTO := openapi.AuthorRequestDTO{
		Id:   authorId,
		Name: "necessary-to-comply-validation",
	}

	author, err := ac.converter.APIToDomain(authorRequestDTO)
	if err != nil {
		return ac.customError(httpServerCtx, ErrGetAuthor, err)
	}

	result, err := ac.service.Get(author)
	if err != nil {
		return ac.customError(httpServerCtx, ErrGetAuthor, err)
	}

	return httpServerCtx.JSONPretty(http.StatusOK, ac.converter.DomainToAPI(result), indent)
}

func (ac *AuthorController) GetAuthors(httpServerCtx server.HttpServerCtx) error {
	authors := ac.service.GetAll()

	return httpServerCtx.JSONPretty(http.StatusOK, ac.converter.DomainsToAPI(authors), indent)
}

func (ac *AuthorController) UpdateAuthor(httpServerCtx server.HttpServerCtx) error {
	authorRequestDTO := &openapi.AuthorRequestDTO{}
	err := httpServerCtx.Bind(authorRequestDTO)
	if err != nil {
		return err
	}

	author, err := ac.converter.APIToDomain(*authorRequestDTO)
	if err != nil {
		return ac.customError(httpServerCtx, ErrUpdateAuthor, err)
	}

	err = ac.service.Update(author)
	if err != nil {
		return ac.customError(httpServerCtx, ErrUpdateAuthor, err)
	}

	return httpServerCtx.JSON(http.StatusOK, ac.converter.DomainToAPI(author))
}

func (ac *AuthorController) DeleteAuthor(httpServerCtx server.HttpServerCtx, authorId string) error {
	authorRequestDTO := openapi.AuthorRequestDTO{
		Id:   authorId,
		Name: "necessary-to-comply-validation",
	}

	author, err := ac.converter.APIToDomain(authorRequestDTO)
	if err != nil {
		return ac.customError(httpServerCtx, ErrDeleteAuthor, err)
	}

	err = ac.service.Delete(author)
	if err != nil {
		return ac.customError(httpServerCtx, ErrDeleteAuthor, err)
	}

	return httpServerCtx.JSON(http.StatusNoContent, ac.converter.DomainToAPI(author))
}

func (ac *AuthorController) customError(httpServerCtx server.HttpServerCtx, message, err error) error {
	m := message.Error()
	e := err.Error()

	return httpServerCtx.JSON(http.StatusBadRequest, openapi.ErrorResponseDTO{
		Message: &m,
		Error:   &e,
	})
}
