package domain

import (
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	// Initialization:

	// Execution:
	user, err := GetUser(0)

	// Validation:

	// Checks if user is empty
	assert.Nil(t, user, "we were not expecting a user with id 0")

	// Check if not nil
	assert.NotNil(t, err)

	// Checks if there is a "Not Found" status code
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)

	// Checks if there is a "Not Found" status code
	assert.EqualValues(t, "not_found", err.Code)

	// Checks if there is an empty error message
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "demby", user.FirstName)
	assert.EqualValues(t, "abella", user.LastName)
}