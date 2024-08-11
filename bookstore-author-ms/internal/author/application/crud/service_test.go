package crud

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"bookstore/bookstore-author-ms/internal/author/domain/model"
)

func TestAuthorService_Create(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		author        model.Author
		expectedError error
	}{
		"Author created": {
			author:        model.Author{Id: "valid-id", BookId: "valid-book-id", Name: "Erik St. Martin"},
			expectedError: nil,
		},
		"No author created by invalid id": {
			author:        model.Author{Id: "invalid-id", BookId: "valid-book-id", Name: "Erik St. Martin"},
			expectedError: errors.New("error description"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Setup
			t.Parallel()

			mockRepository := new(mockRepository)
			mockRepository.On("Create", tc.author).Return(tc.expectedError).Once()

			// Given
			service := AuthorService{
				repository: mockRepository,
			}

			// When
			err := service.Create(tc.author)

			// Then
			if tc.expectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError.Error())
			}

			mockRepository.AssertExpectations(t)
		})
	}
}

func TestAuthorService_Get(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		author         model.Author
		expectedAuthor model.Author
		expectedError  error
	}{
		"Get an author": {
			author:         model.Author{Id: "valid-id", BookId: "valid-book-id", Name: "Erik St. Martin"},
			expectedAuthor: model.Author{Id: "valid-id", BookId: "valid-book-id", Name: "Erik St. Martin"},
			expectedError:  nil,
		},
		"No author by invalid id": {
			author:         model.Author{Id: "invalid-id", BookId: "valid-book-id", Name: "Erik St. Martin"},
			expectedAuthor: model.Author{},
			expectedError:  errors.New("error description"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Setup
			t.Parallel()

			mockRepository := new(mockRepository)
			mockRepository.On("Get", tc.author).Return(tc.expectedAuthor, tc.expectedError).Once()

			// Given
			service := NewAuthorService(mockRepository)

			// When
			author, err := service.Get(tc.author)

			// Then
			if tc.expectedError == nil {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedAuthor, author)
			} else {
				assert.Equal(t, tc.expectedError, err)
			}

			mockRepository.AssertExpectations(t)
		})
	}
}

func TestAuthorService_GetAll(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		authors         []model.Author
		expectedAuthors []model.Author
	}{
		"No authors": {
			authors:         []model.Author{},
			expectedAuthors: []model.Author{},
		},
		"Get an author": {
			authors: []model.Author{
				{Id: "valid-id", Name: "Erik St. Martin"},
			},
			expectedAuthors: []model.Author{
				{Id: "valid-id", Name: "Erik St. Martin"},
			},
		},
		"Get multiple authors": {
			authors: []model.Author{
				{Id: "valid-id-1", Name: "Erik St. Martin"},
				{Id: "valid-id-2", Name: "William Kennedy"},
			},
			expectedAuthors: []model.Author{
				{Id: "valid-id-1", Name: "Erik St. Martin"},
				{Id: "valid-id-2", Name: "William Kennedy"},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Setup
			t.Parallel()

			mockRepository := new(mockRepository)
			mockRepository.On("GetAll").Return(tc.expectedAuthors).Once()

			// Given
			service := NewAuthorService(mockRepository)

			// When
			authors := service.GetAll()

			// Then
			assert.EqualValues(t, tc.expectedAuthors, authors)
			mockRepository.AssertExpectations(t)
		})
	}
}

func TestAuthorService_Update(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		author        model.Author
		expectedError error
	}{
		"Author updated": {
			author:        model.Author{Id: "valid-id", BookId: "valid-book-id", Name: "Erik St. Martin"},
			expectedError: nil,
		},
		"No author updated by invalid id": {
			author:        model.Author{Id: "invalid-id", BookId: "valid-book-id", Name: "Erik St. Martin"},
			expectedError: errors.New("error description"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Setup
			t.Parallel()

			mockRepository := new(mockRepository)
			mockRepository.On("Update", tc.author).Return(tc.expectedError).Once()

			// Given
			service := AuthorService{
				repository: mockRepository,
			}

			// When
			err := service.Update(tc.author)

			// Then
			if tc.expectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError.Error())
			}

			mockRepository.AssertExpectations(t)
		})
	}
}

func TestAuthorService_Delete(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		author        model.Author
		expectedError error
	}{
		"Author deleted": {
			author:        model.Author{Id: "valid-id", BookId: "valid-book-id", Name: "Erik St. Martin"},
			expectedError: nil,
		},
		"No author deleted by invalid id": {
			author:        model.Author{Id: "invalid-id", BookId: "valid-book-id", Name: "Erik St. Martin"},
			expectedError: errors.New("error description"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Setup
			t.Parallel()

			mockRepository := new(mockRepository)
			mockRepository.On("Delete", tc.author).Return(tc.expectedError).Once()

			// Given
			service := AuthorService{
				repository: mockRepository,
			}

			// When
			err := service.Delete(tc.author)

			// Then
			if tc.expectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError.Error())
			}

			mockRepository.AssertExpectations(t)
		})
	}
}
