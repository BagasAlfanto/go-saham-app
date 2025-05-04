package helpers

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

const fileStoragePath = "data/dataUser.csv"

// func init() {
// 	if _, err := os.Stat(FilePath); os.IsNotExist(err) {
// 		os.Mkdir(FilePath, 0755)
// 	}
// }

func FileExists(filename string) bool {
	_, err := os.Stat(fmt.Sprintf("%s/%s", fileStoragePath, filename))
	return !os.IsNotExist(err)
}

func ReadFile(filename string) ([][]string, error) {
	file, err := os.Open(fmt.Sprintf("%s/%s", fileStoragePath, filename))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var records [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}
