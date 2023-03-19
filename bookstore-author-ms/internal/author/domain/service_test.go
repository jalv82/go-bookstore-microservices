package domain

import (
	"testing"

	"bookstore/bookstore-author-ms/internal/author/domain/model"
	"bookstore/bookstore-author-ms/internal/author/infrastructure"
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
	newAuthor := model.Author{
		Id:   uuid.NewString(),
		Name: "Dummy-Name",
	}
	mockRepository.EXPECT().Create(newAuthor).Times(1).Return(nil)
	service := NewAuthorService(mockRepository)

	// When
	err := service.Create(newAuthor)

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
	authorId := uuid.NewString()
	authorToFind := model.Author{
		Id: authorId,
	}
	authorReturn := model.Author{
		Id:   authorId,
		Name: "Dummy-Name",
	}
	mockRepository.EXPECT().Get(authorToFind).Times(1).Return(authorReturn, nil)
	service := NewAuthorService(mockRepository)

	// When
	result, err := service.Get(authorToFind)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, authorToFind.Id, result.Id)
	assert.NotEmpty(t, authorReturn.Name)
}

func TestUpdate(t *testing.T) {
	// Setup
	t.Parallel()
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	// Given
	mockRepository := infrastructure.NewMockRepository(mockController)
	existingAuthor := model.Author{
		Id:   uuid.NewString(),
		Name: "Dummy-Name",
	}
	mockRepository.EXPECT().Update(existingAuthor).Times(1).Return(nil)
	service := NewAuthorService(mockRepository)

	// When
	err := service.Update(existingAuthor)

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
	existingAuthor := model.Author{
		Id:   uuid.NewString(),
		Name: "Dummy-Name",
	}
	mockRepository.EXPECT().Delete(existingAuthor).Times(1).Return(nil)
	service := NewAuthorService(mockRepository)

	// When
	err := service.Delete(existingAuthor)

	// Then
	assert.NoError(t, err)
}
