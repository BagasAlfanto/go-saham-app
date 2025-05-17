package user

import (
	"saham-app/helpers"
	"saham-app/model/transaction"
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
}

var Users []User

var UserLogin *User

/*
 * Load daftar akun dari database
 *
 */
func init() {
	if helpers.FileExists("user.json") {
		content, err := helpers.ReadFile("user.json")
		if err != nil {
			panic(err)
		}

		helpers.LoadFromJSON(content, &Users)

	} else {
		Users = []User{
			{1, "bagas", "bagas", 1000000},
			{2, "lahra", "lahra", 1000000},
		}
		content, err := helpers.SaveToJSON(Users)

		if err != nil {
			panic(err)
		}

		err = helpers.UpdateFile("user.json", content)

		if err != nil {
			panic(err)
		}
	}

}

/*
 * Validasi password user
 *
 */
func CheckPassword(username, password string) (bool, string) {
	for i := range Users {
		if Users[i].Username == username {
			if Users[i].Password == password {
				UserLogin = &Users[i]
				return true, "✅ Login berhasil"
			}
			return false, "❌ Password salah"
		}
	}
	return false, "❌ Username tidak ditemukan"
}

/*
 * Membuat ID baru untuk user
 *
 */
func CreateID() int {
	for _, user := range Users {
		if user.ID > len(Users) {
			return user.ID + 1
		}
	}
	return len(Users) + 1
}

/*
 * Mengimpan data user ke database
 *
 */
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

/*
 * Load username & saldo by user account
 *
 */
func GetUser() User {
	Username := UserLogin.Username
	Saldo := UserLogin.Saldo
	return User{Username: Username, Saldo: Saldo}
}

/*
 * Cek apakah sudah ada username yang sama
 *
 */
func IsUsernameExist(username string) bool {
	for _, user := range Users {
		if user.Username == username {
			return true
		}
	}
	return false
}

/*
 * Menyimpan user baru
 *
 */
func SaveUsers() {
	content, err := helpers.SaveToJSON(Users)
	if err != nil {
		panic(err)
	}

	err = helpers.UpdateFile("user.json", content)
	if err != nil {
		panic(err)
	}
}

/*
 * Load daftar saham by user
 *
 */
func GetPortfolio(userID int) map[string]struct {
	TotalLot   int
	TotalModal int
} {
	result := make(map[string]struct {
		TotalLot   int
		TotalModal int
	})

	for _, t := range transaction.Transactions {
		if t.UserID == userID {
			val := result[t.NamaPerusahaan]
			val.TotalLot += t.JumlahLot
			val.TotalModal += t.Total
			result[t.NamaPerusahaan] = val
		}
	}

	return result
}
