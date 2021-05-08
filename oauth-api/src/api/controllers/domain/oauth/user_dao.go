package oauth

import "github.com/dembygenesis/go-rest-industry-standard/src/api/utils/errors"

const (
	getUserByUsernameAndPassword = "SELECT id, username FORM users WHERE username = ? AND password = ?"
)

var (
	users = map[string]*User{
		"fede": &User{
			Id:       123,
			Username: "fede",
		},
	}
)

func GetUserByUsernameAndPassword(username string, password string) (*User, errors.ApiError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundError("no user found with given parameters")
	}
	return user, nil
}
