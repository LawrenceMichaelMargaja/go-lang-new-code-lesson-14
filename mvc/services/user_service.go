package services

import (
	"github.com/dembygenesis/go-rest-industry-standard/mvc/domain"
	"github.com/dembygenesis/go-rest-industry-standard/mvc/utils"
)

type userService struct {

}

var (
	// Create a public variable to expose the methods
	// This is how we define different artifacts
	// Terminology review: Are artifacts merely public variables containing related methods
	// for a service?
	UserService userService
)

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}