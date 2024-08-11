package api_restful

import (
	"github.com/stretchr/testify/mock"

	"bookstore/bookstore-author-ms/internal/author/domain/model"
)

type MockAuthorService struct {
	mock.Mock
}

func (m *MockAuthorService) Create(a model.Author) error {
	args := m.Called(a)
	return args.Error(0)
}

func (m *MockAuthorService) Get(a model.Author) (model.Author, error) {
	args := m.Called(a)
	return args.Get(0).(model.Author), args.Error(1)
}

func (m *MockAuthorService) GetAll() []model.Author {
	args := m.Called()
	return args.Get(0).([]model.Author)
}

func (m *MockAuthorService) Update(author model.Author) error {
	args := m.Called(author)
	return args.Error(0)
}

func (m *MockAuthorService) Delete(author model.Author) error {
	args := m.Called(author)
	return args.Error(0)
}
