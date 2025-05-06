package controllers

import (
	"fmt"
	"saham-app/handlers/auth"
	"saham-app/helpers"
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
	var username, password, confirmPassword string
	isRegistered := false

	helpers.ClearScreen()
	for !isRegistered {
		fmt.Println("====== Register ======")

		fmt.Print("Username : ")
		fmt.Scan(&username)

		fmt.Print("Password : ")
		fmt.Scan(&password)

		fmt.Print("Ulangi Password : ")
		fmt.Scan(&confirmPassword)

		helpers.ClearScreen()
		for isCorrect := false; !isCorrect; {
			if password == confirmPassword {
				isCorrect = true
			} else {
				fmt.Println("Password tidak sama, silakan coba lagi")
				fmt.Print("Password : ")
				fmt.Scan(&password)

				fmt.Print("Ulangi Password : ")
				fmt.Scan(&confirmPassword)
			}
		}

		user := auth.Register(username, password)
		if user != nil {
			fmt.Println("Register berhasil")
			isRegistered = true
		} else {
			fmt.Println("Username sudah terdaftar, silakan coba lagi")
		}

	}

	return isRegistered
}
