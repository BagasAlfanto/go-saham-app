package helpers

import (
	"fmt"
)

var messages string

/*
 * Clear screen terminal
 *
 */
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

/*
 * Menampilkan menu auth
 *
 */
func DisplayAuthMenu() {
	message := GetMessageValue()

	if message != "" {
		ShowMessages()
		ClearMessage()
	}
	fmt.Println("===== Authenticated =====")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("3. Exit")
	fmt.Println("=========================")
}

/*
 * Menampilkan menu utama
 *
 */
func DisplayMainMenu(username string) {
	message := GetMessageValue()

	if message != "" {
		ShowMessages()
		ClearMessage()
	}
	fmt.Printf("Hai, %s selamat datang!\n", username)
	fmt.Println("===== Main Menu =====")
	fmt.Println("1. Lihat Daftar Saham")
	fmt.Println("2. Beli Saham")
	fmt.Println("3. Jual Saham")
	fmt.Println("4. Lihat Profile")
	fmt.Println("5. Logout")
	fmt.Println("======================")
}

/*
 * Menampilkan menu saham
 *
 */
func DisplaySaham() {
	fmt.Println("======= Saham =======")
	fmt.Println("1. Lihat Semua Saham")
	fmt.Println("2. Cari Saham")
	fmt.Println("3. Urutkan dari harga terrendah")
	fmt.Println("4. Urutkan dari harga tertinggi")
	fmt.Println("5. Kembali Ke Menu Utama")
	fmt.Println("======================")
}

/*
 * Menampilkan tabel daftar saham
 *
 */
func DisplayShowSaham() {
	fmt.Println("===============================================================")
	fmt.Println("|----------------------- Daftar Saham ------------------------|")
	fmt.Println("===============================================================")
	fmt.Printf("| %-15s %-30s %-12s |\n", "Kode Saham", "Nama Perusahaan", "Harga Saham")
	fmt.Println("===============================================================")
}

/*
 * Menampilkan konfirmasi
 *
 */
func ConfirmationScreen() {
	var back string
	for back != "ya" {
		fmt.Println("Kembali ke menu utama? (ya)")
		fmt.Scan(&back)
	}
}

/*
 * Menampilkan dan mendapatkan message
 *
 */
func GetMessages(message string) {
	messages = message
}

func ShowMessages() {
	fmt.Println(messages)
}

func GetMessageValue() string {
	return messages
}

func ClearMessage() {
	messages = ""
}
