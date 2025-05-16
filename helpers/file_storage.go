package helpers

import (
	"fmt"
	"os"
)

/*
 * Set lokasi penyimpanan file json
 *
 */
const fileStoragePath = "./storage"

/*
 * Membuat folder storage jika belum ada
 *
 */
func init() {
	if _, err := os.Stat(fileStoragePath); os.IsNotExist(err) {
		os.Mkdir(fileStoragePath, 0755)
	}
}

/*
 * Mengecek apakah file json ada pada direkttori
 *
 */
func FileExists(filename string) bool {
	_, err := os.Stat(fmt.Sprintf("%s/%s", fileStoragePath, filename))
	return !os.IsNotExist(err)
}

/*
 * Memperbarui file json pada direktori
 *
 */
func UpdateFile(filename string, data []byte) error {
	fullPath := fmt.Sprintf("%s/%s", fileStoragePath, filename)

	if !FileExists(filename) {
		fmt.Println("File belum ada, akan dibuat baru.")
	}

	return os.WriteFile(fullPath, data, 0644)
}

/*
 * Membaca file json pada direktori
 *
 */
func ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(fmt.Sprintf("%s/%s", fileStoragePath, filename))
}
