package domain

import (
	"fmt"
	"github.com/dembygenesis/go-rest-industry-standard/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: &User{Id: 1, FirstName: "demby", LastName: "demby@gg.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
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
