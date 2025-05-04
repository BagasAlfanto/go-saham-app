package user

import (
	"encoding/json"
	"saham-app/helpers"
)

type User struct {
	ID       int
	Username string
	Password string
}

var Users []User

func init() {
	if helpers.FileExists("data/users.csv") {
		content, err := helpers.ReadFile("data/users.csv")

		if err != nil {
			panic(err)
		}

		jsonContent, err := json.Marshal(content)
		if err != nil {
			panic(err)
		}
		helpers.LoadFromJSON(jsonContent, &Users)

	} else {
		Users = []User{
			{ID: 1, Username: "admin", Password: "admin"},
			{ID: 2, Username: "user", Password: "user"},
		}
	}
}

func CheckPassword(username, password string) (bool, string) {
	for _, user := range Users {
		if user.Username == username {
			if user.Password == password {
				return true, "Login berhasil"
			} else {
				return false, "Password salah"
			}
		}
	}
	return false, "Username tidak ditemukan"
}

func CreateID() int {
	for _, user := range Users {
		if user.ID > len(Users) {
			return user.ID + 1
		}
	}
	return len(Users) + 1
}

func Register(username, password string) bool {
	for _, user := range Users {
		if user.Username == username {
			return false
		}
	}
	newUser := User{ID: CreateID(), Username: username, Password: password}
	Users = append(Users, newUser)
	return true
}
