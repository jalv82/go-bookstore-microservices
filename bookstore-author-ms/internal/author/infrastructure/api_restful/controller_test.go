package api_restful

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"bookstore/bookstore-author-ms/internal/author/domain/model"
	"bookstore/bookstore-author-ms/internal/author/infrastructure/api_restful/openapi"
	"bookstore/internal/commons"
)

func TestCreateAuthor(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		request      string
		setupMocks   func(mockService *MockAuthorService, mockConverter *MockAuthorAPIConverter) AuthorController
		expectedCode int
		expectedBody string
	}{
		"Author created": {
			request: `{
	"id": "valid-id",
	"name": "William Kennedy", 
	"bookId": "valid-book-id"
}`,
			setupMocks: func(mockService *MockAuthorService, mockConverter *MockAuthorAPIConverter) AuthorController {
				modelAuthor := model.Author{
					Id:     "valid-id",
					Name:   "William Kennedy",
					BookId: "valid-book-id",
				}

				responseAuthor := openapi.AuthorResponseDTO{
					Id:     "valid-id",
					Name:   "William Kennedy",
					BookId: commons.ValueOfPointer("valid-book-id"),
				}

				mockConverter.On("APIToDomain", mock.AnythingOfType("openapi.AuthorRequestDTO")).Return(modelAuthor, nil).Once()
				mockService.On("Create", mock.AnythingOfType("model.Author")).Return(nil).Once()
				mockConverter.On("DomainToAPI", mock.AnythingOfType("model.Author")).Return(responseAuthor).Once()

				return AuthorController{service: mockService, converter: mockConverter}
			},
			expectedCode: http.StatusCreated,
			expectedBody: `{"bookId":"valid-book-id","id":"valid-id","name":"William Kennedy"}
`,
		},
		"No author created by invalid id": {
			request: `{
	"id": "invalid-id",
	"name": "William Kennedy", 
	"bookId": "valid-book-id"
}`,
			setupMocks: func(mockService *MockAuthorService, mockConverter *MockAuthorAPIConverter) AuthorController {
				modelAuthor := model.Author{
					Id:     "invalid-id",
					Name:   "William Kennedy",
					BookId: "valid-book-id",
				}

				mockConverter.On("APIToDomain", mock.AnythingOfType("openapi.AuthorRequestDTO")).Return(modelAuthor, nil).Once()
				mockService.On("Create", mock.AnythingOfType("model.Author")).Return(errors.New("author not found")).Once()

				return AuthorController{service: mockService, converter: mockConverter}
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"author not found","message":"error creating Author"}
`,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//Setup
			t.Parallel()

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tc.request))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e := echo.New()
			ectx := e.NewContext(req, rec)

			mockService := new(MockAuthorService)
			mockConverter := new(MockAuthorAPIConverter)

			// Given
			controller := tc.setupMocks(mockService, mockConverter)

			// When
			err := controller.CreateAuthor(ectx)

			// Then
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedCode, rec.Code)
			assert.Equal(t, tc.expectedBody, rec.Body.String())
			mockConverter.AssertExpectations(t)
			mockService.AssertExpectations(t)
		})
	}
}

func TestAuthorController_GetAuthor(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		authorId     string
		setupMocks   func(mockService *MockAuthorService, mockConverter *MockAuthorAPIConverter) AuthorController
		expectedCode int
		expectedBody string
	}{
		"Get an author": {
			authorId: "valid-id",
			setupMocks: func(mockService *MockAuthorService, mockConverter *MockAuthorAPIConverter) AuthorController {
				modelAuthor := model.Author{
					Id:     "valid-id",
					Name:   "William Kennedy",
					BookId: "valid-book-id",
				}

				responseAuthor := openapi.AuthorResponseDTO{
					Id:     "valid-id",
					Name:   "William Kennedy",
					BookId: commons.ValueOfPointer("valid-book-id"),
				}

				mockConverter.On("APIToDomain", mock.AnythingOfType("openapi.AuthorRequestDTO")).Return(modelAuthor, nil).Once()
				mockService.On("Get", mock.AnythingOfType("model.Author")).Return(modelAuthor, nil).Once()
				mockConverter.On("DomainToAPI", mock.AnythingOfType("model.Author")).Return(responseAuthor).Once()

				return AuthorController{service: mockService, converter: mockConverter}
			},
			expectedCode: http.StatusOK,
			expectedBody: `{
  "bookId": "valid-book-id",
  "id": "valid-id",
  "name": "William Kennedy"
}
`,
		},
		"No author by nonexistent id": {
			authorId: "nonexistent-id",
			setupMocks: func(mockService *MockAuthorService, mockConverter *MockAuthorAPIConverter) AuthorController {
				modelAuthor := model.Author{
					Id:     "nonexistent-id",
					Name:   "William Kennedy",
					BookId: "valid-book-id",
				}

				mockConverter.On("APIToDomain", mock.AnythingOfType("openapi.AuthorRequestDTO")).Return(modelAuthor, nil).Once()
				mockService.On("Get", mock.AnythingOfType("model.Author")).Return(model.Author{}, errors.New("author not found")).Once()

				return AuthorController{service: mockService, converter: mockConverter}
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"author not found","message":"error getting Author"}
`,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Setup
			t.Parallel()

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			ectx := e.NewContext(req, rec)

			mockService := new(MockAuthorService)
			mockConverter := new(MockAuthorAPIConverter)

			// Given
			controller := tc.setupMocks(mockService, mockConverter)

			// When
			err := controller.GetAuthor(ectx, tc.authorId)

			// Then
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedCode, rec.Code)
			assert.Equal(t, tc.expectedBody, rec.Body.String())
			mockService.AssertExpectations(t)
			mockConverter.AssertExpectations(t)
		})
	}
}

func TestAuthorController_GetAuthors(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		setupMocks   func(mockService *MockAuthorService, mockConverter *MockAuthorAPIConverter) AuthorController
		expectedCode int
		expectedBody string
	}{
		"No authors": {
			setupMocks: func(mockService *MockAuthorService, mockConverter *MockAuthorAPIConverter) AuthorController {
				expectedAuthors := []model.Author{}
				responseAuthors := []openapi.AuthorResponseDTO{}

				mockService.On("GetAll").Return(expectedAuthors, nil).Once()
				mockConverter.On("DomainsToAPI", mock.AnythingOfType("[]model.Author")).Return(responseAuthors).Once()

				return AuthorController{service: mockService, converter: mockConverter}
			},
			expectedCode: http.StatusOK,
			expectedBody: `[]
`,
		},
		"Get an author": {
			setupMocks: func(mockService *MockAuthorService, mockConverter *MockAuthorAPIConverter) AuthorController {
				expectedAuthors := []model.Author{
					{
						Id:     "valid-id",
						Name:   "William Kennedy",
						BookId: "valid-book-id",
					},
				}

				responseAuthors := []openapi.AuthorResponseDTO{
					{
						Id:     "valid-id",
						Name:   "William Kennedy",
						BookId: commons.ValueOfPointer("valid-book-id"),
					},
				}

				mockService.On("GetAll").Return(expectedAuthors, nil).Once()
				mockConverter.On("DomainsToAPI", mock.AnythingOfType("[]model.Author")).Return(responseAuthors).Once()

				return AuthorController{service: mockService, converter: mockConverter}
			},
			expectedCode: http.StatusOK,
			expectedBody: `[
  {
    "bookId": "valid-book-id",
    "id": "valid-id",
    "name": "William Kennedy"
  }
]
`,
		},
		"Get multiple authors": {
			setupMocks: func(mockService *MockAuthorService, mockConverter *MockAuthorAPIConverter) AuthorController {
				expectedAuthors := []model.Author{
					{
						Id:     "valid-id-1",
						Name:   "Erik St. Martin",
						BookId: "valid-book-id-1",
					},
					{
						Id:     "valid-id-2",
						Name:   "William Kennedy",
						BookId: "valid-book-id-2",
					},
				}

				responseAuthors := []openapi.AuthorResponseDTO{
					{
						Id:     "valid-id-1",
						Name:   "Erik St. Martin",
						BookId: commons.ValueOfPointer("valid-book-id-1"),
					},
					{
						Id:     "valid-id-2",
						Name:   "William Kennedy",
						BookId: commons.ValueOfPointer("valid-book-id-2"),
					},
				}

				mockService.On("GetAll").Return(expectedAuthors, nil).Once()
				mockConverter.On("DomainsToAPI", mock.AnythingOfType("[]model.Author")).Return(responseAuthors).Once()

				return AuthorController{service: mockService, converter: mockConverter}
			},
			expectedCode: http.StatusOK,
			expectedBody: `[
  {
    "bookId": "valid-book-id-1",
    "id": "valid-id-1",
    "name": "Erik St. Martin"
  },
  {
    "bookId": "valid-book-id-2",
    "id": "valid-id-2",
    "name": "William Kennedy"
  }
]
`,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Setup
			t.Parallel()

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			ectx := e.NewContext(req, rec)

			mockService := new(MockAuthorService)
			mockConverter := new(MockAuthorAPIConverter)

			// Given
			controller := tc.setupMocks(mockService, mockConverter)

			// When
			err := controller.GetAuthors(ectx)

			// Then
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedCode, rec.Code)
			assert.Equal(t, tc.expectedBody, rec.Body.String())
			mockService.AssertExpectations(t)
			mockConverter.AssertExpectations(t)
		})
	}
}

func TestUpdateAuthor(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		request      string
		setupMocks   func(mockService *MockAuthorService, mockConverter *MockAuthorAPIConverter) AuthorController
		expectedCode int
		expectedBody string
	}{
		"Author updated": {
			request: `{
	"id": "valid-id",
	"name": "William Kennedy", 
	"bookId": "valid-book-id"
}`,
			setupMocks: func(mockService *MockAuthorService, mockConverter *MockAuthorAPIConverter) AuthorController {
				modelAuthor := model.Author{
					Id:     "valid-id",
					Name:   "William Kennedy",
					BookId: "valid-book-id",
				}

				responseAuthor := openapi.AuthorResponseDTO{
					Id:     "valid-id",
					Name:   "William Kennedy",
					BookId: commons.ValueOfPointer("valid-book-id"),
				}

				mockConverter.On("APIToDomain", mock.AnythingOfType("openapi.AuthorRequestDTO")).Return(modelAuthor, nil).Once()
				mockService.On("Update", mock.AnythingOfType("model.Author")).Return(nil).Once()
				mockConverter.On("DomainToAPI", mock.AnythingOfType("model.Author")).Return(responseAuthor).Once()

				return AuthorController{service: mockService, converter: mockConverter}
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"bookId":"valid-book-id","id":"valid-id","name":"William Kennedy"}
`,
		},
		"No author updated by invalid id": {
			request: `{
	"id": "invalid-id",
	"name": "William Kennedy", 
	"bookId": "valid-book-id"
}`,
			setupMocks: func(mockService *MockAuthorService, mockConverter *MockAuthorAPIConverter) AuthorController {
				modelAuthor := model.Author{
					Id:     "invalid-id",
					Name:   "William Kennedy",
					BookId: "valid-book-id",
				}

				mockConverter.On("APIToDomain", mock.AnythingOfType("openapi.AuthorRequestDTO")).Return(modelAuthor, nil).Once()
				mockService.On("Update", mock.AnythingOfType("model.Author")).Return(errors.New("author not found")).Once()

				return AuthorController{service: mockService, converter: mockConverter}
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"author not found","message":"error updating Author"}
`,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//Setup
			t.Parallel()

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tc.request))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e := echo.New()
			ectx := e.NewContext(req, rec)

			mockService := new(MockAuthorService)
			mockConverter := new(MockAuthorAPIConverter)

			// Given
			controller := tc.setupMocks(mockService, mockConverter)

			// When
			err := controller.UpdateAuthor(ectx)

			// Then
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedCode, rec.Code)
			assert.Equal(t, tc.expectedBody, rec.Body.String())
			mockService.AssertExpectations(t)
			mockConverter.AssertExpectations(t)
		})
	}
}

func TestDeleteAuthor(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		authorId     string
		setupMocks   func(mockService *MockAuthorService, mockConverter *MockAuthorAPIConverter) AuthorController
		expectedCode int
		expectedBody string
	}{
		"Author deleted": {
			authorId: "valid-id",
			setupMocks: func(mockService *MockAuthorService, mockConverter *MockAuthorAPIConverter) AuthorController {
				modelAuthor := model.Author{
					Id: "valid-id",
				}

				responseAuthor := openapi.AuthorResponseDTO{
					Id:     "valid-id",
					Name:   "William Kennedy",
					BookId: commons.ValueOfPointer("valid-book-id"),
				}

				mockConverter.On("APIToDomain", mock.AnythingOfType("openapi.AuthorRequestDTO")).Return(modelAuthor, nil).Once()
				mockService.On("Delete", mock.AnythingOfType("model.Author")).Return(nil).Once()
				mockConverter.On("DomainToAPI", mock.AnythingOfType("model.Author")).Return(responseAuthor).Once()

				return AuthorController{service: mockService, converter: mockConverter}
			},
			expectedCode: http.StatusNoContent,
			expectedBody: `{"bookId":"valid-book-id","id":"valid-id","name":"William Kennedy"}
`,
		},
		"No author deleted by nonexistent id": {
			authorId: "nonexistent-id",
			setupMocks: func(mockService *MockAuthorService, mockConverter *MockAuthorAPIConverter) AuthorController {
				modelAuthor := model.Author{
					Id: "nonexistent-id",
				}

				mockConverter.On("APIToDomain", mock.AnythingOfType("openapi.AuthorRequestDTO")).Return(modelAuthor, nil).Once()
				mockService.On("Delete", mock.AnythingOfType("model.Author")).Return(errors.New("author not found")).Once()

				return AuthorController{service: mockService, converter: mockConverter}
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"author not found","message":"error deleting Author"}
`,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//Setup
			t.Parallel()

			req := httptest.NewRequest(http.MethodPost, "/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e := echo.New()
			ectx := e.NewContext(req, rec)

			mockService := new(MockAuthorService)
			mockConverter := new(MockAuthorAPIConverter)

			// Given
			controller := tc.setupMocks(mockService, mockConverter)

			// When
			err := controller.DeleteAuthor(ectx, tc.authorId)

			// Then
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedCode, rec.Code)
			assert.Equal(t, tc.expectedBody, rec.Body.String())
			mockService.AssertExpectations(t)
			mockConverter.AssertExpectations(t)
		})
	}
}
