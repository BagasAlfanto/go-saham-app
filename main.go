package main

import (
	"fmt"
	"saham-app/controllers"
	"saham-app/helpers"
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
		helpers.ClearScreen()
		helpers.DisplayMainMenu()

		for choice < 1 || choice > 3 {
			fmt.Print("Masukan Pilihan Anda : ")
			fmt.Scan(&choice)
		}

		switch choice {
		case 1:
		case 2:
		case 3:
		case 4:
		}
	}

	fmt.Println("=== Program Selesai ===")

}
