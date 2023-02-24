package database

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var (
	expectedBook = &Book{
		Id:       "default-id",
		AuthorId: "default-author-id",
		Title:    "default-title",
	}

	expectedError = errors.New("with this query the database have done BOOM")
)

func Test_Book_SQLClient_Create_OK(t *testing.T) {
	// Setup
	t.Parallel()
	mock, sqlClient := NewMockBookSqlClient()
	query := `INSERT INTO "books" ("id","author_id","title") VALUES ($1,$2,$3)`

	// Given
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(expectedBook.Id, expectedBook.AuthorId, expectedBook.Title).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// When
	err := sqlClient.Create(expectedBook)

	// Then
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func Test_Book_SQLClient_Create_Return_Error(t *testing.T) {
	// Setup
	t.Parallel()
	mock, sqlClient := NewMockBookSqlClient()
	query := `INSERT INTO "books" ("id","author_id","title") VALUES ($1,$2,$3)`

	// Given
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(expectedBook.Id, expectedBook.AuthorId, expectedBook.Title).
		WillReturnError(expectedError)
	mock.ExpectRollback()

	// When
	err := sqlClient.Create(expectedBook)

	// Then
	assert.ErrorIs(t, err, expectedError)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func Test_Book_SQLClient_Get_OK(t *testing.T) {
	// Setup
	t.Parallel()
	mock, sqlClient := NewMockBookSqlClient()
	query := `SELECT * FROM "books"`
	columns := []string{"id", "author_id", "title"}

	// Given
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(expectedBook.Id, expectedBook.AuthorId, expectedBook.Title))

	// When
	actualBook, err := sqlClient.Get(expectedBook)

	// Then
	assert.NoError(t, err)
	assert.NotEmpty(t, actualBook)
	assert.Equal(t, actualBook, expectedBook)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func Test_Book_SQLClient_Get_Return_Err(t *testing.T) {
	// Setup
	t.Parallel()
	mock, sqlClient := NewMockBookSqlClient()
	query := `SELECT * FROM "books"`

	// Given
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnError(expectedError)

	// When
	actualBook, err := sqlClient.Get(expectedBook)

	// Then
	assert.ErrorIs(t, err, expectedError)
	assert.Empty(t, actualBook)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func Test_Book_SQLClient_Update_OK(t *testing.T) {
	// Setup
	t.Parallel()
	mock, sqlClient := NewMockBookSqlClient()
	query := `UPDATE "books" SET "author_id"=$1,"title"=$2 WHERE "id" = $3`

	// Given
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(expectedBook.AuthorId, expectedBook.Title, expectedBook.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// When
	err := sqlClient.Update(expectedBook)

	// Then
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func Test_Book_SQLClient_Update_Return_Error(t *testing.T) {
	// Setup
	t.Parallel()
	mock, sqlClient := NewMockBookSqlClient()
	query := `UPDATE "books" SET "author_id"=$1,"title"=$2 WHERE "id" = $3`

	// Given
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WillReturnError(expectedError)
	mock.ExpectRollback()

	// When
	err := sqlClient.Update(expectedBook)

	// Then
	assert.ErrorIs(t, err, expectedError)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func Test_Book_SQLClient_Delete_OK(t *testing.T) {
	// Setup
	t.Parallel()
	mock, sqlClient := NewMockBookSqlClient()
	query := `DELETE FROM "books" WHERE "books"."id" = $1`

	// Given
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(expectedBook.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// When
	err := sqlClient.Delete(expectedBook)

	// Then
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func Test_Book_SQLClient_Delete_Return_Error(t *testing.T) {
	// Setup
	t.Parallel()
	mock, sqlClient := NewMockBookSqlClient()
	query := `DELETE FROM "books" WHERE "books"."id" = $1`

	// Given
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WillReturnError(expectedError)
	mock.ExpectRollback()

	// When
	err := sqlClient.Delete(expectedBook)

	// Then
	assert.ErrorIs(t, err, expectedError)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
