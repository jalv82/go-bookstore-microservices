package acceptance_tests

import (
	"context"
	"errors"

	"bookstore/bookstore-author-ms/internal/author/domain"
	"bookstore/bookstore-author-ms/internal/author/domain/model"
	"bookstore/bookstore-author-ms/internal/author/infrastructure"
	"bookstore/bookstore-author-ms/internal/author/infrastructure/database"
	"bookstore/internal/tests"
	"github.com/cucumber/godog"
)

type AuthorContext struct {
	author  *model.Author
	service *domain.AuthorService
}

func (ac *AuthorContext) aDefaultAuthor() (err error) {
	ac.author = &model.Author{
		Id:     "default-id",
		BookId: "default-book-id",
		Name:   "default-name",
	}

	return
}

func (ac *AuthorContext) theAuthorIdIs(id string) (err error) {
	ac.author.Id = id
	return
}

func (ac *AuthorContext) theAuthorNameIs(name string) (err error) {
	ac.author.Name = name
	return
}

func (ac *AuthorContext) theAuthorBookIdIs(id string) (err error) {
	ac.author.BookId = id
	return
}

func (ac *AuthorContext) theAuthorIsCreatedInDb() (err error) {
	err = ac.service.Create(*ac.author)
	if err != nil {
		return
	}

	return
}

func (ac *AuthorContext) theAuthorIsUpdatedInDb() (err error) {
	err = ac.service.Update(*ac.author)
	if err != nil {
		return
	}

	return
}

func (ac *AuthorContext) theAuthorIsObtainedInDb() (err error) {
	author, err := ac.service.Get(*ac.author)
	if err != nil {
		return
	}

	ac.author = &author

	return
}

func (ac *AuthorContext) theAuthorIsDeletedInDb() (err error) {
	err = ac.service.Delete(*ac.author)
	if err != nil {
		return
	}

	return
}

func (ac *AuthorContext) theAuthorIdIsNotExitsInDb() (err error) {
	_, err = ac.service.Get(*ac.author)
	if err == nil {
		return errors.New("the author still exits")
	}

	return nil
}

func (ac *AuthorContext) theAuthorIdReadInDbIs(id string) (err error) {
	author := &model.Author{
		Id: id,
	}

	_, err = ac.service.Get(*author)
	if err != nil {
		return
	}

	return
}

func (ac *AuthorContext) theAuthorNameReadInDbIs(name string) (err error) {
	author := &model.Author{
		Name: name,
	}

	_, err = ac.service.Get(*author)
	if err != nil {
		return
	}

	return
}

func (ac *AuthorContext) theAuthorBookIdReadInDbIs(bookId string) (err error) {
	author := &model.Author{
		BookId: bookId,
	}

	_, err = ac.service.Get(*author)
	if err != nil {
		return
	}

	return
}

func InitializeAuthorScenario(sctx *godog.ScenarioContext) {
	authorSQLClient := database.NewAuthorSqlClient(tests.GetTestEnvironment().DbConfig)
	authorSQLConverter := database.NewAuthorSQLConverter()
	repository := infrastructure.NewRepository(*authorSQLClient, *authorSQLConverter)

	actx := &AuthorContext{
		service: domain.NewAuthorService(repository),
	}

	sctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		_, _ = tests.GetTestEnvironment().Db.Exec("TRUNCATE TABLE authors")
		return ctx, nil
	})

	sctx.Step(`^a default author`, actx.aDefaultAuthor)

	sctx.Step(`^the author id is "([^"]*)"$`, actx.theAuthorIdIs)
	sctx.Step(`^the author name is "([^"]*)"$`, actx.theAuthorNameIs)
	sctx.Step(`^the author book-id is "([^"]*)"$`, actx.theAuthorBookIdIs)

	sctx.Step(`^the author is created in author-db$`, actx.theAuthorIsCreatedInDb)
	sctx.Step(`^the author is updated in author-db$`, actx.theAuthorIsUpdatedInDb)
	sctx.Step(`^the author is obtained in author-db$`, actx.theAuthorIsObtainedInDb)
	sctx.Step(`^the author is deleted in author-db$`, actx.theAuthorIsDeletedInDb)

	sctx.Step(`^the author id read in author-db is "([^"]*)"$`, actx.theAuthorIdReadInDbIs)
	sctx.Step(`^the author name read in author-db is "([^"]*)"$`, actx.theAuthorNameReadInDbIs)
	sctx.Step(`^the author book-id read in author-db is "([^"]*)"$`, actx.theAuthorBookIdReadInDbIs)

	sctx.Step(`^the author id "([^"]*)" isn't exists in author-db$`, actx.theAuthorIdIsNotExitsInDb)
}
