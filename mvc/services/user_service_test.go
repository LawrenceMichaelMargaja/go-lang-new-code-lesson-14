package services

import (
	"fmt"
	"github.com/dembygenesis/go-rest-industry-standard/mvc/domain"
	"github.com/dembygenesis/go-rest-industry-standard/mvc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

/**
This part of the tutorial is where you are introduced to Mocking your service in order to make it testable.
Steps:
	1. Override the artifacts you will be calling (even if indirectly through another service)
		with another struct that implements the same interface
	2. Make your struct implement the method that you want to overwrite, and ensure
		the function has the same signatures
	3. Declare the function (used dynamically and still implements the same signature as the artifact function aimed
		to be mocked)
	4. Just modify the function dynamically per Test function to make it suit the logic the test function is trying to
		achieve
*/

var (
	userDaoMock usersDaoMock

	// Step 3:
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	// Step 1:
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

// Step 2:
func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	fmt.Println("This function is overridden")
	return getUserFunction(userId)
}

func TestUserNotFoundInDatabase(t *testing.T) {
	// Step 4:
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    "user 0 was not found",
			StatusCode: http.StatusNotFound,
		}
	}

	// This is where I missed some details. The user service is calling "GetUser"
	// but what is easily missed is that UserService calls the domain.UserDao
	// which was modified here to be another struct, but still complies with the same interface.
	// So, it's GetUser method was overridden xoxo
	user, err := UserService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	// Step 4:
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{Id: 123}, nil
	}

	user, err := UserService.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}
