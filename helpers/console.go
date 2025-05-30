package helpers

import (
	"bufio"
	"fmt"
	"os"
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
	message := GetMessageValue()

	if message != "" {
		ShowMessages()
		ClearMessage()
	}
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
 * Format nominal mata uang Indonesia
 *
 */
func NominalFormat(nominal int) string {
  return fmt.Sprintf("Rp%s", formatRibuan(nominal))
}

/*
 * Convert format ribuan 
 *
 */
func formatRibuan(n int) string {
  s := fmt.Sprintf("%d", n)
  nLen := len(s)
  if nLen <= 3 {
    return s
  }
  var result []byte
  mod := nLen % 3
  if mod > 0 {
    result = append(result, s[:mod]...)
    if nLen > mod {
      result = append(result, '.')
    }
  }
  for i := mod; i < nLen; i += 3 {
    result = append(result, s[i:i+3]...)
    if i+3 < nLen {
      result = append(result, '.')
    }
  }
  return string(result)
}

/*
 * Menampilkan konfirmasi dengan menekan Enter
 *
 */
func ConfirmationScreen() {
	fmt.Println("Tekan Enter untuk kembali ke menu utama...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	reader.ReadString('\n')
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
