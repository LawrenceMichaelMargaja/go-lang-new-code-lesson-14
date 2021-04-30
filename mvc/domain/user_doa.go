package domain

import (
	"fmt"
	"github.com/dembygenesis/go-rest-industry-standard/mvc/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		123: &User{Id: 123, FirstName: "demby", LastName: "abella"},
	}


	// UserDao userDao

	// This was previously of type "userDao" but, now is an interface type.
	// This is fine as long as it has all the methods and properties required by the
	// interface.
	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDao struct {

}


// !!! This is how I prevent database access when running tests.
// I make an interface
type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("we're accessing the database")

	user := users[userId]
	if user == nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("user %v was not found", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	return user, nil
}
