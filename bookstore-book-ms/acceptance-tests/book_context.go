package acceptance_tests

import (
	"context"
	"errors"

	"bookstore/bookstore-book-ms/internal/book/domain"
	"bookstore/bookstore-book-ms/internal/book/domain/model"
	"bookstore/bookstore-book-ms/internal/book/infrastructure"
	"bookstore/bookstore-book-ms/internal/book/infrastructure/database"
	"bookstore/internal/tests"
	"github.com/cucumber/godog"
)

type BookContext struct {
	book    *model.Book
	service *domain.BookService
	repo    *infrastructure.BookRepository
}

func (bc *BookContext) aDefaultBook() (err error) {
	bc.book = &model.Book{
		Id:       "default-id",
		AuthorId: "default-author-id",
		Title:    "default-title",
	}

	return
}

func (bc *BookContext) theBookIdIs(id string) (err error) {
	bc.book.Id = id
	return
}

func (bc *BookContext) theBookTitleIs(title string) (err error) {
	bc.book.Title = title
	return
}

func (bc *BookContext) theBookAuthorIdIs(id string) (err error) {
	bc.book.AuthorId = id
	return
}

func (bc *BookContext) theBookIsCreatedInDb() (err error) {
	err = bc.service.Create(*bc.book)
	if err != nil {
		return
	}

	return
}

func (bc *BookContext) theBookIsUpdatedInDb() (err error) {
	err = bc.service.Update(*bc.book)
	if err != nil {
		return
	}

	return
}

func (bc *BookContext) theBookIsObtainedInDb() (err error) {
	book, err := bc.service.Get(*bc.book)
	if err != nil {
		return
	}

	bc.book = &book

	return
}

func (bc *BookContext) theBookIsDeletedInDb() (err error) {
	err = bc.service.Delete(*bc.book)
	if err != nil {
		return
	}

	return
}

func (bc *BookContext) theBookIdIsNotExitsInDb() (err error) {
	_, err = bc.service.Get(*bc.book)
	if err == nil {
		return errors.New("the book still exits")
	}

	return nil
}

func (bc *BookContext) theBookIdReadInDbIs(id string) (err error) {
	book := &model.Book{
		Id: id,
	}

	_, err = bc.service.Get(*book)
	if err != nil {
		return
	}

	return
}

func (bc *BookContext) theBookTitleReadInDbIs(title string) (err error) {
	book := &model.Book{
		Title: title,
	}

	_, err = bc.service.Get(*book)
	if err != nil {
		return
	}

	return
}

func (bc *BookContext) theBookAuthorIdReadInDbIs(authorId string) (err error) {
	book := &model.Book{
		AuthorId: authorId,
	}

	_, err = bc.service.Get(*book)
	if err != nil {
		return
	}

	return
}

func InitializeBookScenario(sctx *godog.ScenarioContext) {
	bookSQLClient := database.NewBookSqlClient(tests.GetTestEnvironment().DbConfig)
	bookSQLConverter := database.NewBookSQLConverter()
	repository := infrastructure.NewRepository(*bookSQLClient, *bookSQLConverter)

	bctx := &BookContext{
		service: domain.NewBookService(repository),
	}

	sctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		_, _ = tests.GetTestEnvironment().Db.Exec("TRUNCATE TABLE books")
		return ctx, nil
	})

	sctx.Step(`^a default book$`, bctx.aDefaultBook)

	sctx.Step(`^the book id is "([^"]*)"$`, bctx.theBookIdIs)
	sctx.Step(`^the book title is "([^"]*)"$`, bctx.theBookTitleIs)
	sctx.Step(`^the book author-id is "([^"]*)"$`, bctx.theBookAuthorIdIs)

	sctx.Step(`^the book is created in book-db$`, bctx.theBookIsCreatedInDb)
	sctx.Step(`^the book is updated in book-db$`, bctx.theBookIsUpdatedInDb)
	sctx.Step(`^the book is obtained in book-db$`, bctx.theBookIsObtainedInDb)
	sctx.Step(`^the book is deleted in book-db$`, bctx.theBookIsDeletedInDb)

	sctx.Step(`^the book id read in book-db is "([^"]*)"$`, bctx.theBookIdReadInDbIs)
	sctx.Step(`^the book title read in book-db is "([^"]*)"$`, bctx.theBookTitleReadInDbIs)
	sctx.Step(`^the book author-id read in book-db is "([^"]*)"$`, bctx.theBookAuthorIdReadInDbIs)

	sctx.Step(`^the book id "([^"]*)" isn't exists in book-db$`, bctx.theBookIdIsNotExitsInDb)
}
