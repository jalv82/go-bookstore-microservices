package acceptance_tests

import (
	"bookstore/internal/core/author/domain"
	"bookstore/internal/core/author/domain/model"
	"github.com/cucumber/godog"
)

type AuthorContext struct {
	author  *model.Author
	service *domain.AuthorService
}

func (a *AuthorContext) aDefaultAuthor() (err error) {
	a.author = &model.Author{
		Id:     "default-id",
		BookId: "default-book-id",
		Name:   "default-name",
	}

	return
}
func (a *AuthorContext) theAuthorNameIs(name string) error {
	a.author.Name = name
	return nil
}

func (a *AuthorContext) theAuthorBookIdIs(id string) error {
	a.author.BookId = id
	return nil
}

func InitializeAuthorScenario(sctx *godog.ScenarioContext) {
	actx := &AuthorContext{}

	sctx.Step(`^a default author`, actx.aDefaultAuthor)
	sctx.Step(`^the author name is "([^"]*)"$`, actx.theAuthorNameIs)
	sctx.Step(`^the author book-id is "([^"]*)"$`, actx.theAuthorBookIdIs)

	// TODO - For implements the next steps, is necessary to use TestContainers
	// sctx.Step(`^the author is registered in book-db$`, actx.theAuthorIsRegisteredInDb)
	// sctx.Step(`^the author id read in author-db is "([^"]*)"$`, actx.theAuthorIdReadInDbIs)
	// sctx.Step(`^the author name read in author-db is "([^"]*)"$`, actx.theAuthorNameReadInDbIs)
	// sctx.Step(`^the author book-id read in author-db is "([^"]*)"$`, actx.theAuthorBookIdReadInDbIs)
}
