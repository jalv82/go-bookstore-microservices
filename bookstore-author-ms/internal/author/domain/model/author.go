package model

import (
	"github.com/invopop/validation"
	"github.com/invopop/validation/is"
	"github.com/rs/zerolog/log"
)

type Author struct {
	Id     string
	BookId string
	Name   string
}

func NewAuthor(id, bookId, name string) (Author, error) {
	author := Author{
		Id:     id,
		BookId: bookId,
		Name:   name,
	}

	err := author.validate()
	if err != nil {
		log.Error().Str("id", author.Id).Msg(err.Error())

		return Author{}, err
	}

	return author, nil
}

func (a Author) validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Id, validation.Required, is.UUIDv4),
		validation.Field(&a.BookId, validation.When(a.BookId != "", validation.Required, is.UUIDv4)),
		validation.Field(&a.Name, validation.Required, validation.Length(2, 64)),
	)
}
