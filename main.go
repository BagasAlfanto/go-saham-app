// ================================================
// Aplikasi Simulasi Saham (saham-app)
// ------------------------------------------------
// saham-app merupakan aplikasi simulasi saham
// sederhana yang dapat memudahkan pengguna
// untuk belajar tentang dunia saham.
// Apliaksi ini dibuat menggunakan bahasa Go
// dan hanya bisa dijalankan didalam terminal
// ------------------------------------------------
// Version     : 1.0
// Date        : 2025-04-28
// License     : Unlicense
// OS          : Windows
// Language    : Go
// ================================================

package main

import (
	"fmt"
	"saham-app/controllers"
	"saham-app/helpers"
	"saham-app/model/saham"
	"saham-app/model/user"
)

func main() {
	var choice int
	isLogged := false

	for !isLogged {
		helpers.ClearScreen()
		helpers.DisplayAuthMenu()

		choice = 0
		for choice < 1 || choice > 3 {
			fmt.Print("Masukan Pilihan Anda : ")
			fmt.Scan(&choice)
		}

		helpers.ClearScreen()

		switch choice {
		case 1:
			isLogged = controllers.Login()
		case 2:
			controllers.Register()
		case 3:
			isLogged = true
		}

		if choice != 3 {
			choice = 0
		}
	}

	if choice == 3 {
		helpers.ClearScreen()
		fmt.Println("Terima kasih telah menggunakan saham-app. Selamat tinggal!")
		return
	}

	for isRunning := false; !isRunning; {
		accountUsed := user.GetUser()

		helpers.ClearScreen()

		helpers.DisplayMainMenu(accountUsed.Username)

		for choice < 1 || choice > 5 {
			fmt.Print("Masukan Pilihan Anda : ")
			fmt.Scan(&choice)
		}
		helpers.ClearScreen()

		switch choice {
		case 1:
			controllers.SahamMenu()
		case 2:
			controllers.BuyingSaham()
		case 3:
			controllers.SellSaham()
		case 4:
			controllers.ShowProfile()

		case 5:
			isRunning = true
			saham.SaveSaham()
			fmt.Println("Terimakasih sudah menggunakan saham-app, selamat tinggal..")
		}

		if choice != 5 {
			choice = 0
		}
		
	}

}
