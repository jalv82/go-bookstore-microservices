package crud

import (
	"github.com/stretchr/testify/mock"

	"bookstore/bookstore-author-ms/internal/author/domain/model"
)

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) Create(author model.Author) error {
	args := m.Called(author)
	return args.Error(0)
}

func (m *mockRepository) Get(author model.Author) (model.Author, error) {
	args := m.Called(author)
	return args.Get(0).(model.Author), args.Error(1)
}

func (m *mockRepository) GetAll() []model.Author {
	args := m.Called()
	return args.Get(0).([]model.Author)
}

func (m *mockRepository) Update(author model.Author) error {
	args := m.Called(author)
	return args.Error(0)
}

func (m *mockRepository) Delete(author model.Author) error {
	args := m.Called(author)
	return args.Error(0)
}
