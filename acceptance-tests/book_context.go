package acceptance_tests

import (
	"bookstore/internal/core/book/domain"
	"bookstore/internal/core/book/domain/model"
	"github.com/cucumber/godog"
)

type BookContext struct {
	book    *model.Book
	service *domain.BookService
}

func (a *BookContext) aDefaultBook() (err error) {
	a.book = &model.Book{
		Id:       "default-id",
		AuthorId: "default-author-id",
		Title:    "default-title",
	}

	return
}
func (a *BookContext) theBookTitleIs(title string) error {
	a.book.Title = title
	return nil
}

func (a *BookContext) theBookAuthorIdIs(id string) error {
	a.book.AuthorId = id
	return nil
}

func InitializeBookScenario(sctx *godog.ScenarioContext) {
	bctx := &BookContext{}

	sctx.Step(`^a default book$`, bctx.aDefaultBook)
	sctx.Step(`^the book title is "([^"]*)"$`, bctx.theBookTitleIs)
	sctx.Step(`^the book author-id is "([^"]*)"$`, bctx.theBookAuthorIdIs)

	// TODO - For implements the next steps, is necessary to use TestContainers
	// sctx.Step(`^the book is registered in book-db$`, bctx.theBookIsRegisteredInDb)
	// sctx.Step(`^the book id read in book-db is "([^"]*)"$`, bctx.theBookIdReadInDbIs)
	// sctx.Step(`^the book title read in book-db is "([^"]*)"$`, bctx.theBookTitleReadInDbIs)
	// sctx.Step(`^the book author-id read in book-db is "([^"]*)"$`, bctx.theBookAuthorIdReadInDbIs)
}
