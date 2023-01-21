package domain

import (
	"testing"

	"bookstore/internal/core/author/domain/model"
	"bookstore/internal/core/author/infrastructure"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRepository := infrastructure.NewMockRepository(mockController)
	newAuthor := model.Author{
		Id:   uuid.NewString(),
		Name: "Dummy-Name",
	}
	mockRepository.EXPECT().Create(newAuthor).Times(1).Return(nil)
	service := NewAuthorService(mockRepository)

	err := service.Create(newAuthor)

	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

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

	result, err := service.Get(authorToFind)

	assert.NoError(t, err)
	assert.Equal(t, authorToFind.Id, result.Id)
	assert.NotEmpty(t, authorReturn.Name)
}

func TestUpdate(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRepository := infrastructure.NewMockRepository(mockController)
	existingAuthor := model.Author{
		Id:   uuid.NewString(),
		Name: "Dummy-Name",
	}
	mockRepository.EXPECT().Update(existingAuthor).Times(1).Return(nil)
	service := NewAuthorService(mockRepository)

	err := service.Update(existingAuthor)

	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRepository := infrastructure.NewMockRepository(mockController)
	existingAuthor := model.Author{
		Id:   uuid.NewString(),
		Name: "Dummy-Name",
	}
	mockRepository.EXPECT().Delete(existingAuthor).Times(1).Return(nil)
	service := NewAuthorService(mockRepository)

	err := service.Delete(existingAuthor)

	assert.NoError(t, err)
}
