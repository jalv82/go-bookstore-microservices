package database

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var (
	expectedAuthor = &Author{
		Id:     "default-id",
		BookId: "default-book-id",
		Name:   "default-name",
	}

	expectedError = errors.New("with this query the database have done BOOM")
)

func Test_Author_SQLClient_Create_OK(t *testing.T) {
	// Setup
	t.Parallel()
	mock, sqlClient := NewMockAuthorSqlClient()
	query := `INSERT INTO "authors" ("id","book_id","name") VALUES ($1,$2,$3)`

	// Given
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(expectedAuthor.Id, expectedAuthor.BookId, expectedAuthor.Name).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// When
	err := sqlClient.Create(expectedAuthor)

	// Then
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func Test_Author_SQLClient_Create_Return_Error(t *testing.T) {
	// Setup
	t.Parallel()
	mock, sqlClient := NewMockAuthorSqlClient()
	query := `INSERT INTO "authors" ("id","book_id","name") VALUES ($1,$2,$3)`

	// Given
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(expectedAuthor.Id, expectedAuthor.BookId, expectedAuthor.Name).
		WillReturnError(expectedError)
	mock.ExpectRollback()

	// When
	err := sqlClient.Create(expectedAuthor)

	// Then
	assert.ErrorIs(t, err, expectedError)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func Test_Author_SQLClient_Get_OK(t *testing.T) {
	// Setup
	t.Parallel()
	mock, sqlClient := NewMockAuthorSqlClient()
	query := `SELECT * FROM "authors"`
	columns := []string{"id", "book_id", "name"}

	// Given
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(expectedAuthor.Id, expectedAuthor.BookId, expectedAuthor.Name))

	// When
	actualAuthor, err := sqlClient.Get(expectedAuthor)

	// Then
	assert.NoError(t, err)
	assert.NotEmpty(t, actualAuthor)
	assert.Equal(t, actualAuthor, expectedAuthor)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func Test_Author_SQLClient_Get_Return_Err(t *testing.T) {
	// Setup
	t.Parallel()
	mock, sqlClient := NewMockAuthorSqlClient()
	query := `SELECT * FROM "authors"`

	// Given
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnError(expectedError)

	// When
	actualAuthor, err := sqlClient.Get(expectedAuthor)

	// Then
	assert.ErrorIs(t, err, expectedError)
	assert.Empty(t, actualAuthor)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func Test_Author_SQLClient_Update_OK(t *testing.T) {
	// Setup
	t.Parallel()
	mock, sqlClient := NewMockAuthorSqlClient()
	query := `UPDATE "authors" SET "book_id"=$1,"name"=$2 WHERE "id" = $3`

	// Given
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(expectedAuthor.BookId, expectedAuthor.Name, expectedAuthor.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// When
	err := sqlClient.Update(expectedAuthor)

	// Then
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func Test_Author_SQLClient_Update_Return_Error(t *testing.T) {
	// Setup
	t.Parallel()
	mock, sqlClient := NewMockAuthorSqlClient()
	query := `UPDATE "authors" SET "book_id"=$1,"name"=$2 WHERE "id" = $3`

	// Given
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WillReturnError(expectedError)
	mock.ExpectRollback()

	// When
	err := sqlClient.Update(expectedAuthor)

	// Then
	assert.ErrorIs(t, err, expectedError)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func Test_Author_SQLClient_Delete_OK(t *testing.T) {
	// Setup
	t.Parallel()
	mock, sqlClient := NewMockAuthorSqlClient()
	query := `DELETE FROM "authors" WHERE "authors"."id" = $1`

	// Given
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(expectedAuthor.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// When
	err := sqlClient.Delete(expectedAuthor)

	// Then
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func Test_Author_SQLClient_Delete_Return_Error(t *testing.T) {
	// Setup
	t.Parallel()
	mock, sqlClient := NewMockAuthorSqlClient()
	query := `DELETE FROM "authors" WHERE "authors"."id" = $1`

	// Given
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WillReturnError(expectedError)
	mock.ExpectRollback()

	// When
	err := sqlClient.Delete(expectedAuthor)

	// Then
	assert.ErrorIs(t, err, expectedError)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
