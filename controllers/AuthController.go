package controllers

import (
	"fmt"
	"saham-app/handlers/auth"
	"saham-app/helpers"
)

/*
 *Menangani proses login
 *
 */
func Login() bool {
	var username, password string
	isLogged := false

	for !isLogged {
		for i := 1; i <= 3; i++ {
			fmt.Println("====== Login ======")
			fmt.Print("Username : ")
			fmt.Scan(&username)

			fmt.Print("Password : ")
			fmt.Scan(&password)

			helpers.ClearScreen()
			if auth.Authentic(username, password) {
				isLogged = true
				break

			} else {
				helpers.ShowMessages()
			}
		}
		if !isLogged {
			helpers.GetMessages("Jika anda lupa akun anda silahkan buat akun kembali")
			break
		}
	}

	return isLogged
}

/*
 *Menangani proses Register
 *
 */
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
				fmt.Println("❌ Password tidak sama, silakan coba lagi")
				fmt.Print("Password : ")
				fmt.Scan(&password)

				fmt.Print("Ulangi Password : ")
				fmt.Scan(&confirmPassword)
			}
			helpers.ClearScreen()
		}

		user := auth.Register(username, password)
		if user != nil {
			helpers.GetMessages("✅ Register Berhasil")
			isRegistered = true
		} else {
			helpers.GetMessages("❌ Username sudah terdaftar, silahkan coba lagi")
			break
		}
	}

	return isRegistered
}
