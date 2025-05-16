package transaction

import (
	"fmt"
	"saham-app/helpers"
)

/*
 * Struct transaction saham
 *
 */
type Transaction struct {
	TransactionID  int
	UserID         int
	NamaPerusahaan string
	JumlahLot      int
	HargaPerLembar int
	Total          int
	Tipe           string
}

var Transactions []Transaction

/*
 * Load daftar transaksi dari database
 *
 */
func init() {
	if helpers.FileExists("transactions.json") {
		content, err := helpers.ReadFile("transactions.json")
		if err != nil {
			panic(err)
		}

		helpers.LoadFromJSON(content, &Transactions)
	} else {
		Transactions = []Transaction{}
		content, err := helpers.SaveToJSON(Transactions)

		if err != nil {
			panic(err)
		}

		err = helpers.UpdateFile("transactions.json", content)

		if err != nil {
			panic(err)
		}
	}
}

/*
 * Menyimpan history transaksi ke database
 *
 */
func AppendToHistory(t Transaction) {
	var histories []Transaction

	// Load history.json
	if helpers.FileExists("history.json") {
		content, err := helpers.ReadFile("history.json")
		if err == nil {
			helpers.LoadFromJSON(content, &histories)
		}
	}

	// Tambahkan transaksi baru
	t.TransactionID = len(histories) + 1
	histories = append(histories, t)

	content, err := helpers.SaveToJSON(histories)
	if err != nil {
		fmt.Println("❌ Gagal menyimpan ke history:", err)
		return
	}

	err = helpers.UpdateFile("history.json", content)
	if err != nil {
		fmt.Println("❌ Gagal update file history.json:", err)
	}
}

/*
 * Menyimpan transaksi beli saham ke database
 *
 */
func SaveTransactions() error {
	content, err := helpers.SaveToJSON(Transactions)
	if err != nil {
		return err
	}
	return helpers.UpdateFile("transactions.json", content)
}

func InsertTransaction(t Transaction) {
	t.TransactionID = len(Transactions) + 1
	Transactions = append(Transactions, t)
	if err := SaveTransactions(); err != nil {
		fmt.Println("❌ Gagal menyimpan transaksi:", err)
	}
}

/*
 * Load daftar transaksi by user 
 *
 */
func GetTransactionsByUserID(userID int) []Transaction {
	var result []Transaction
	for _, t := range Transactions {
		if t.UserID == userID {
			result = append(result, t)
		}
	}
	return result
}

/*
 * Load daftar history by user
 *
 */
func GetHistoryByUserID(userID int) []Transaction {
	var histories []Transaction

	if !helpers.FileExists("history.json") {
		return histories
	}

	content, err := helpers.ReadFile("history.json")
	if err != nil {
		return histories
	}

	helpers.LoadFromJSON(content, &histories)

	var result []Transaction
	for _, t := range histories {
		if t.UserID == userID {
			result = append(result, t)
		}
	}

	return result
}
