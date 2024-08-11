package domain

import "bookstore/bookstore-author-ms/internal/author/domain/model"

type CRUDPort interface {
	Create(author model.Author) error
	Get(author model.Author) (model.Author, error)
	GetAll() []model.Author
	Update(author model.Author) error
	Delete(author model.Author) error
}
