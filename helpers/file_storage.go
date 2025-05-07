package helpers

import (
	"fmt"
	"os"
)

const fileStoragePath = "./storage"

func FileExists(filename string) bool {
	_, err := os.Stat(fmt.Sprintf("%s/%s", fileStoragePath, filename))
	return !os.IsNotExist(err)
}

func UpdateFile(filename string, data []byte) error {
	fullPath := fmt.Sprintf("%s/%s", fileStoragePath, filename)

	if !FileExists(filename) {
		fmt.Println("File belum ada, akan dibuat baru.")
	}

	return os.WriteFile(fullPath, data, 0666)
}

func ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(fmt.Sprintf("%s/%s", fileStoragePath, filename))
}
