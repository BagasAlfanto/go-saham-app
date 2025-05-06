package user

import (
	"saham-app/helpers"
)

type User struct {
	ID       int
	Username string
	Password string
	Saldo    int
}

var Users []User

func init() {
	if helpers.FileExists("user.json") {
		content, err := helpers.ReadFile("user.json")
		if err != nil {
			panic(err)
		}

		helpers.LoadFromJSON(content, &Users)

	} else {
		Users = []User{
			{ID: 1, Username: "admin", Password: "admin", Saldo: 1000000},
			{ID: 2, Username: "user", Password: "user", Saldo: 1000000},
		}
		content, err := helpers.SaveToJSON(Users)

		if err != nil {
			panic(err)
		}

		err = helpers.SaveFile("user.json", content)

		if err != nil {
			panic(err)
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

func InsertUser(user User) {
	user.ID = CreateID()
	Users = append(Users, user)

	content, err := helpers.SaveToJSON(Users)
	if err != nil {
		panic(err)
	}

	err = helpers.UpdateFile("user.json", content)
	if err != nil {
		panic(err)
	}
}

// func Register(username, password string) bool {
// 	for _, user := range Users {
// 		if user.Username == username {
// 			return false
// 		}
// 	}
// 	newUser := User{ID: CreateID(), Username: username, Password: password}
// 	Users = append(Users, newUser)
// 	return true
// }
