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

func SaveFile(filename string, data []byte) error {
	if FileExists(filename) {
		return fmt.Errorf("file already exists")
	}

	return os.WriteFile(fmt.Sprintf("%s/%s", fileStoragePath, filename), data, 0666)
}

func UpdateFile(filename string, data []byte) error {
	if !FileExists(filename) {
		return fmt.Errorf("file does not exist")
	}

	return os.WriteFile(fmt.Sprintf("%s/%s", fileStoragePath, filename), data, 0666)
}

func ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(fmt.Sprintf("%s/%s", fileStoragePath, filename))
}
