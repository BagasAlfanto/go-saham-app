package helpers

import (
	"fmt"
)

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

func DisplayMainMenu(username string) {
	fmt.Println("===== Main Menu =====")
	fmt.Printf("Hai, %s selamat datang!\n", username)
	fmt.Println("1. Lihat Daftar Saham")
	fmt.Println("2. Beli Saham")
	fmt.Println("3. Lihat Profile")
	fmt.Println("4. Logout")
	fmt.Println("======================")
}

func DisplaySaham(){
	fmt.Println("======= Saham =======")
	fmt.Println("1. Lihat Semua Saham")
	fmt.Println("2. Cari Saham")
	fmt.Println("3. Urutkan dari harga terrendah")
	fmt.Println("4. Urutkan dari harga tertinggi")
	fmt.Println("5. Kembali Ke Menu Utama")
	fmt.Println("======================")
}

func DisplayShowSaham(){
	fmt.Println("===============================================================")
	fmt.Println("|----------------------- Daftar Saham ------------------------|")
	fmt.Println("===============================================================")
	fmt.Printf("| %-15s %-30s %-12s |\n", "Kode Saham", "Nama Perusahaan", "Harga Saham")
	fmt.Println("===============================================================")
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
