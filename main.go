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
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi")
			helpers.Confirmation()
		}

		if choice != 3 {
			choice = 0
		}
	}

	if choice == 3 {
		helpers.ClearScreen()
		fmt.Println("Terima kasih telah menggunakan Aplikasi Ini. Selamat tinggal!")
		return
	}

	for isRunning := false; !isRunning; {
		accountUsed := user.GetUser()

		helpers.ClearScreen()
		helpers.DisplayMainMenu(accountUsed.Username)

		for choice < 1 || choice > 3 {
			fmt.Print("Masukan Pilihan Anda : ")
			fmt.Scan(&choice)
		}
		helpers.ClearScreen()

		switch choice {
		case 1:
			controllers.SahamMenu()
		case 2:
			saham.UpdatePrice()
		case 3:
			controllers.ShowProfile()
		case 4:
			isRunning = true
		}

		if choice != 4 {
			choice = 0
		}
	}

	fmt.Println("=== Program Selesai ===")

}