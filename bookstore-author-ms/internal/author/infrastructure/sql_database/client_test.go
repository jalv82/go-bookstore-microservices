package sql_database

import (
	"errors"
	"regexp"
	"testing"

	"bookstore/bookstore-author-ms/internal/author/infrastructure/sql_database/model"
	"bookstore/internal/commons"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAuthorSQLClient_Create(t *testing.T) {
	t.Parallel()

	testsCases := map[string]struct {
		authorToCreate model.Author
		setupMocks     func() (sqlmock.Sqlmock, *AuthorSQLClient)
		expectedError  error
	}{
		"Author created": {
			authorToCreate: model.Author{
				Id:     "valid-id",
				Name:   "William Kennedy",
				BookId: "valid-book-id",
			},
			setupMocks: func() (sqlmock.Sqlmock, *AuthorSQLClient) {
				mock, sqlClient := NewMockAuthorSQLClient(commons.NewPostgreSQLClientMock())
				query := `INSERT INTO "authors" ("id","book_id","name") VALUES ($1,$2,$3)`

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs("valid-id", "valid-book-id", "William Kennedy").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()

				return mock, sqlClient
			},
			expectedError: nil,
		},
		"No author created because id already exists": {
			authorToCreate: model.Author{
				Id:     "valid-id",
				Name:   "William Kennedy",
				BookId: "valid-book-id",
			},
			setupMocks: func() (sqlmock.Sqlmock, *AuthorSQLClient) {
				mock, sqlClient := NewMockAuthorSQLClient(commons.NewPostgreSQLClientMock())
				query := `INSERT INTO "authors" ("id","book_id","name") VALUES ($1,$2,$3)`

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs("valid-id", "valid-book-id", "William Kennedy").
					WillReturnError(errors.New("record already exists with id: valid-id"))
				mock.ExpectRollback()

				return mock, sqlClient
			},
			expectedError: errors.New("record already exists with id: valid-id"),
		},
	}

	for name, tc := range testsCases {
		t.Run(name, func(t *testing.T) {
			// Setup
			t.Parallel()

			// Given
			mock, sqlClient := tc.setupMocks()

			// When
			err := sqlClient.Create(&tc.authorToCreate)

			// Then
			assert.Equal(t, err, tc.expectedError)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestAuthorSQLClient_Get(t *testing.T) {
	t.Parallel()

	testsCases := map[string]struct {
		authorToGet   model.Author
		setupMocks    func() (sqlmock.Sqlmock, *AuthorSQLClient)
		expectedError error
	}{
		"Get an author": {
			authorToGet: model.Author{
				Id:     "valid-id",
				Name:   "William Kennedy",
				BookId: "valid-book-id",
			},
			setupMocks: func() (sqlmock.Sqlmock, *AuthorSQLClient) {
				mock, sqlClient := NewMockAuthorSQLClient(commons.NewPostgreSQLClientMock())
				query := `SELECT * FROM "authors" WHERE "authors"."id" = $1 ORDER BY "authors"."id" LIMIT $2`
				columns := []string{"id", "book_id", "name"}

				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs("valid-id", 1).
					WillReturnRows(sqlmock.NewRows(columns).
						AddRow("valid-id", "valid-book-id", "William Kennedy"))

				return mock, sqlClient
			},
			expectedError: nil,
		},
		"No author by nonexistent id": {
			authorToGet: model.Author{
				Id:     "nonexistent-id",
				Name:   "William Kennedy",
				BookId: "valid-book-id",
			},
			setupMocks: func() (sqlmock.Sqlmock, *AuthorSQLClient) {
				mock, sqlClient := NewMockAuthorSQLClient(commons.NewPostgreSQLClientMock())
				query := `SELECT * FROM "authors" WHERE "authors"."id" = $1 ORDER BY "authors"."id" LIMIT $2`

				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs("nonexistent-id", 1).
					WillReturnError(errors.New("description error"))

				return mock, sqlClient
			},
			expectedError: errors.New("description error"),
		},
	}

	for name, tc := range testsCases {
		t.Run(name, func(t *testing.T) {
			// Setup
			t.Parallel()

			// Given
			mock, sqlClient := tc.setupMocks()

			// When
			err := sqlClient.Get(&tc.authorToGet)

			// Then
			assert.Equal(t, err, tc.expectedError)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestAuthorSQLClient_GetAll(t *testing.T) {
	t.Parallel()

	testsCases := map[string]struct {
		setupMocks      func() (sqlmock.Sqlmock, *AuthorSQLClient)
		expectedAuthors []model.Author
	}{
		"No authors": {
			setupMocks: func() (sqlmock.Sqlmock, *AuthorSQLClient) {
				mock, sqlClient := NewMockAuthorSQLClient(commons.NewPostgreSQLClientMock())
				query := `SELECT * FROM "authors"`

				mock.ExpectQuery(regexp.QuoteMeta(query))

				return mock, sqlClient
			},
			expectedAuthors: []model.Author(nil),
		},
		"Get an author": {
			setupMocks: func() (sqlmock.Sqlmock, *AuthorSQLClient) {
				mock, sqlClient := NewMockAuthorSQLClient(commons.NewPostgreSQLClientMock())
				query := `SELECT * FROM "authors"`
				columns := []string{"id", "book_id", "name"}

				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WillReturnRows(sqlmock.NewRows(columns).
						AddRow("valid-id", "valid-book-id", "William Kennedy"))

				return mock, sqlClient
			},
			expectedAuthors: []model.Author{
				{
					Id:     "valid-id",
					Name:   "William Kennedy",
					BookId: "valid-book-id",
				},
			},
		},
		"Get multiple authors": {
			setupMocks: func() (sqlmock.Sqlmock, *AuthorSQLClient) {
				mock, sqlClient := NewMockAuthorSQLClient(commons.NewPostgreSQLClientMock())
				query := `SELECT * FROM "authors"`
				columns := []string{"id", "book_id", "name"}

				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WillReturnRows(sqlmock.NewRows(columns).
						AddRow("valid-id-1", "valid-book-id-1", "Erik St. Martin").
						AddRow("valid-id-2", "valid-book-id-2", "William Kennedy"))

				return mock, sqlClient
			},
			expectedAuthors: []model.Author{
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
			},
		},
	}

	for name, tc := range testsCases {
		t.Run(name, func(t *testing.T) {
			// Setup
			t.Parallel()

			// Given
			mock, sqlClient := tc.setupMocks()

			// When
			actualAuthors := sqlClient.GetAll()

			// Then
			assert.Equal(t, tc.expectedAuthors, actualAuthors)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestAuthorSQLClient_Update(t *testing.T) {
	t.Parallel()

	testsCases := map[string]struct {
		authorToUpdate model.Author
		setupMocks     func() (sqlmock.Sqlmock, *AuthorSQLClient)
		expectedError  error
	}{
		"Author updated": {
			authorToUpdate: model.Author{
				Id:     "valid-id",
				Name:   "William Kennedy",
				BookId: "valid-book-id",
			},
			setupMocks: func() (sqlmock.Sqlmock, *AuthorSQLClient) {
				mock, sqlClient := NewMockAuthorSQLClient(commons.NewPostgreSQLClientMock())

				selectQuery := `SELECT * FROM "authors" WHERE "authors"."id" = $1 ORDER BY "authors"."id" LIMIT $2`
				columns := []string{"id", "book_id", "name"}
				updateQuery := `UPDATE "authors"`

				mock.ExpectQuery(regexp.QuoteMeta(selectQuery)).
					WithArgs("valid-id", 1).
					WillReturnRows(sqlmock.NewRows(columns).
						AddRow("valid-id", "valid-book-id", "William Kennedy"))

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(updateQuery)).
					WithArgs("valid-book-id", "William Kennedy", "valid-id").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()

				return mock, sqlClient
			},
			expectedError: nil,
		},
		"No author updated by invalid id": {
			authorToUpdate: model.Author{
				Id:     "valid-id",
				Name:   "William Kennedy",
				BookId: "valid-book-id",
			},
			setupMocks: func() (sqlmock.Sqlmock, *AuthorSQLClient) {
				mock, sqlClient := NewMockAuthorSQLClient(commons.NewPostgreSQLClientMock())

				selectQuery := `SELECT * FROM "authors" WHERE "authors"."id" = $1 ORDER BY "authors"."id" LIMIT $2`

				mock.ExpectQuery(regexp.QuoteMeta(selectQuery)).
					WithArgs("valid-id", 1).
					WillReturnError(errors.New("record not found with id: valid-id"))

				return mock, sqlClient
			},
			expectedError: errors.New("record not found with id: valid-id"),
		},
	}

	for name, tc := range testsCases {
		t.Run(name, func(t *testing.T) {
			// Setup
			t.Parallel()

			// Given
			mock, sqlClient := tc.setupMocks()

			// When
			err := sqlClient.Update(&tc.authorToUpdate)

			// Then
			assert.Equal(t, err, tc.expectedError)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestAuthorSQLClient_Delete(t *testing.T) {
	t.Parallel()

	testsCases := map[string]struct {
		authorToDelete model.Author
		setupMocks     func() (sqlmock.Sqlmock, *AuthorSQLClient)
		expectedError  error
	}{
		"Author deleted": {
			authorToDelete: model.Author{
				Id:     "valid-id",
				Name:   "William Kennedy",
				BookId: "valid-book-id",
			},
			setupMocks: func() (sqlmock.Sqlmock, *AuthorSQLClient) {
				mock, sqlClient := NewMockAuthorSQLClient(commons.NewPostgreSQLClientMock())
				query := `DELETE FROM "authors" WHERE "authors"."id" = $1`

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs("valid-id").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()

				return mock, sqlClient
			},
			expectedError: nil,
		},
		"No author deleted by nonexistent id": {
			authorToDelete: model.Author{
				Id:     "nonexistent-id",
				Name:   "William Kennedy",
				BookId: "valid-book-id",
			},
			setupMocks: func() (sqlmock.Sqlmock, *AuthorSQLClient) {
				mock, sqlClient := NewMockAuthorSQLClient(commons.NewPostgreSQLClientMock())
				query := `DELETE FROM "authors" WHERE "authors"."id" = $1`

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs("nonexistent-id").
					WillReturnError(errors.New("record not found with id: nonexistent-id"))
				mock.ExpectRollback()

				return mock, sqlClient
			},
			expectedError: errors.New("record not found with id: nonexistent-id"),
		},
	}

	for name, tc := range testsCases {
		t.Run(name, func(t *testing.T) {
			// Setup
			t.Parallel()

			// Given
			mock, sqlClient := tc.setupMocks()

			// When
			err := sqlClient.Delete(&tc.authorToDelete)

			// Then
			assert.Equal(t, err, tc.expectedError)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
