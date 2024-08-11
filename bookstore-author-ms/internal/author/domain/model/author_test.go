package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAuthor(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		id          string
		bookId      string
		name        string
		expectError string
	}{
		"Valid author": {
			id:          "dda36a73-407b-4d30-bc9e-ea912be2875c",
			bookId:      "9fd5dd6d-86cc-4eaa-9bff-54b081650c77",
			name:        "Brian Ketelsen",
			expectError: "",
		},
		"Valid author with empty bookId": {
			id:          "dda36a73-407b-4d30-bc9e-ea912be2875c",
			bookId:      "",
			name:        "Brian Ketelsen",
			expectError: "",
		},
		"Invalid author by missing id": {
			id:          "",
			bookId:      "9fd5dd6d-86cc-4eaa-9bff-54b081650c77",
			name:        "Brian Ketelsen",
			expectError: "Id: cannot be blank.",
		},
		"Invalid author by invalid id": {
			id:          "invalid id",
			bookId:      "9fd5dd6d-86cc-4eaa-9bff-54b081650c77",
			name:        "Brian Ketelsen",
			expectError: "Id: must be a valid UUID v4.",
		},
		"Invalid author by invalid bookId": {
			id:          "dda36a73-407b-4d30-bc9e-ea912be2875c",
			bookId:      "invalid bookId",
			name:        "Brian Ketelsen",
			expectError: "BookId: must be a valid UUID v4.",
		},
		"Invalid author by empty name": {
			id:          "dda36a73-407b-4d30-bc9e-ea912be2875c",
			bookId:      "9fd5dd6d-86cc-4eaa-9bff-54b081650c77",
			name:        "",
			expectError: "Name: cannot be blank.",
		},
		"Invalid author by short name (< 2 characters)": {
			id:          "dda36a73-407b-4d30-bc9e-ea912be2875c",
			bookId:      "9fd5dd6d-86cc-4eaa-9bff-54b081650c77",
			name:        "B",
			expectError: "Name: the length must be between 2 and 64.",
		},
		"Invalid author by long name (> 64 characters)": {
			id:          "dda36a73-407b-4d30-bc9e-ea912be2875c",
			bookId:      "9fd5dd6d-86cc-4eaa-9bff-54b081650c77",
			name:        "Brian Ketelsen + aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			expectError: "Name: the length must be between 2 and 64.",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Setup
			t.Parallel()

			// When
			author, err := NewAuthor(tc.id, tc.bookId, tc.name)

			// Then
			if tc.expectError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.id, author.Id)
				assert.Equal(t, tc.bookId, author.BookId)
				assert.Equal(t, tc.name, author.Name)
			} else {
				assert.Error(t, err)
				assert.Equal(t, Author{}, author)
			}
		})
	}
}
