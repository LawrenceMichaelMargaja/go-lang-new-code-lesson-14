package services

import (
	"github.com/dembygenesis/go-rest-industry-standard/mvc/domain"
	"github.com/dembygenesis/go-rest-industry-standard/mvc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

// Ahhhh... If we are in a test file, Go automatically identifies
// it, and point it up to the interface
type usersDaoMock struct {

}

func (m* usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

// TestGetUserNotFoundInDatabase - tests using a mock function
func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message: "user 0 does not exist",
		}
	}

	user, err := getUserFunction(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 does not exist", err.Message)
}


func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{Id: 123}, nil
	}

	user, err := getUserFunction(123)
	assert.Nil(t,err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}