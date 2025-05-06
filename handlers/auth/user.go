package auth

import (
	"fmt"
	"saham-app/model/user"
)

func Authentic(username, password string) bool {
	success, message := user.CheckPassword(username, password)

	fmt.Println(message)
	return success
}

func Register(username, password string) *user.User {
	userData := user.User{
		ID:       user.CreateID(),
		Username: username,
		Password: password,
		Saldo:    1000000,
	}

	user.InsertUser(userData)
	return &userData
}
