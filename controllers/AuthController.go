package controllers

import (
	"fmt"
	"saham-app/handlers/auth"
	"saham-app/helpers"
	"saham-app/model/user"
)

func Login() bool {
	var username, password string
	isLogged := false

	for !isLogged {
		fmt.Println("====== Login ======")
		fmt.Print("Username : ")
		fmt.Scan(&username)

		fmt.Print("Password : ")
		fmt.Scan(&password)

		helpers.ClearScreen()
		if auth.Authentic(username, password) {
			isLogged = true
		} else {
			fmt.Println("Login gagal, silakan coba lagi")
		}
	}

	return isLogged

}

func Register() bool {
	var username, password, retryPassword string
	isRegistered := false

	helpers.ClearScreen()
	for !isRegistered {
		fmt.Println("====== Register ======")

		fmt.Print("Username : ")
		fmt.Scan(&username)

		fmt.Print("Password : ")
		fmt.Scan(&password)

		fmt.Print("Ulangi Password : ")
		fmt.Scan(&retryPassword)

		helpers.ClearScreen()
		if password != retryPassword {
			fmt.Println("Password tidak sama")
		} else if user.Register(username, password) {
			fmt.Println("Registrasi berhasil")
			isRegistered = true
		} else {
			fmt.Println("Registrasi gagal, username sudah terdaftar")
		}
	}

	return isRegistered
}
