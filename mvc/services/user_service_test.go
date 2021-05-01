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
)

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return &domain.User{}, nil
}

func TestUserNotFoundInDatabase(t *testing.T) {
	user, err := UserService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := UserService.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}