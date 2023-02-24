package domain

import (
	"testing"

	"bookstore/internal/core/book/domain/model"
	"bookstore/internal/core/book/infrastructure"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	// Setup
	t.Parallel()
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	// Given
	mockRepository := infrastructure.NewMockRepository(mockController)
	newBook := model.Book{
		Id:    uuid.NewString(),
		Title: "Dummy-Title",
	}
	mockRepository.EXPECT().Create(newBook).Times(1).Return(nil)
	service := NewBookService(mockRepository)

	// When
	err := service.Create(newBook)

	// Then
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// Setup
	t.Parallel()
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	// Given
	mockRepository := infrastructure.NewMockRepository(mockController)
	bookId := uuid.NewString()
	bookToFind := model.Book{
		Id: bookId,
	}
	bookReturn := model.Book{
		Id:    bookId,
		Title: "Dummy-Title",
	}
	mockRepository.EXPECT().Get(bookToFind).Times(1).Return(bookReturn, nil)
	service := NewBookService(mockRepository)

	// When
	result, err := service.Get(bookToFind)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, bookToFind.Id, result.Id)
	assert.NotEmpty(t, bookReturn.Title)
}

func TestUpdate(t *testing.T) {
	// Setup
	t.Parallel()
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	// Given
	mockRepository := infrastructure.NewMockRepository(mockController)
	existingBook := model.Book{
		Id:    uuid.NewString(),
		Title: "Dummy-Title",
	}
	mockRepository.EXPECT().Update(existingBook).Times(1).Return(nil)
	service := NewBookService(mockRepository)

	// When
	err := service.Update(existingBook)

	// Then
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	// Setup
	t.Parallel()
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	// Given
	mockRepository := infrastructure.NewMockRepository(mockController)
	existingBook := model.Book{
		Id:    uuid.NewString(),
		Title: "Dummy-Title",
	}
	mockRepository.EXPECT().Delete(existingBook).Times(1).Return(nil)
	service := NewBookService(mockRepository)

	// When
	err := service.Delete(existingBook)

	// Then
	assert.NoError(t, err)
}
