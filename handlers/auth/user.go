package auth

import (
	"saham-app/helpers"
	"saham-app/model/user"
)

/*
 * Struct User
 *
 */
type User struct {
	ID       int
	Username string
	Password string
	Saldo    int
	Saham    []string
}

/*
 * Variabel struct user & User account
 *
 */
var Users []user.User
var UserLogin *user.User

/*
 * Authtentic User
 *
 */
func Authentic(username, password string) bool {
	success, message := user.CheckPassword(username, password)

	helpers.GetMessages(message)
	return success
}

/*
 * Register User
 *
 */
func Register(username, password string) *user.User {

	if user.IsUsernameExist(username) {
		return nil
	}
	userData := user.User{
		ID:       user.CreateID(),
		Username: username,
		Password: password,
		Saldo:    1000000,
	}

	user.InsertUser(userData)
	return &userData
}
