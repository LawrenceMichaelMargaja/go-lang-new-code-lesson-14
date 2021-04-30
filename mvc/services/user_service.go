package services

import (
	"github.com/dembygenesis/go-rest-industry-standard/mvc/domain"
	"github.com/dembygenesis/go-rest-industry-standard/mvc/utils"
)

// !!! Seeked architecture knowledge
// This is how you isolate methods in the same package,
// by mapping them onto a struct
type userService struct {

}

// !!! Seeked architecture knowledge
// This is the equivalent of how you expose a class in Go
// Created a public variable
var (
	UserService userService
)

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}