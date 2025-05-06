package helpers

import "fmt"

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func Confirmation() {
	fmt.Println("Tekan tombol enter untuk melanjutkan...")
	fmt.Scanln()
}

func DisplayAuthMenu() {
	fmt.Println("===== Authenticated =====")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("3. Exit")
	fmt.Println("=========================")
}

func DisplayMainMenu() {
	fmt.Println("===== Main Menu =====")
	fmt.Println("1. Lihat Daftar Saham")
	fmt.Println("2. Beli Saham")
	fmt.Println("3. Lihat Profile")
	fmt.Println("4. Logout")
	fmt.Println("======================")
}

func ConfirmationScreen(messages ...string) {
	for _, message := range messages {
		fmt.Println(message)
	}

	if len(messages) > 0 {
		fmt.Println()
	}

	fmt.Println("Tekan tombol enter untuk melanjutkan...")
	fmt.Scanln()
}
