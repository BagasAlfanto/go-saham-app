package selling

import (
	"fmt"
	"saham-app/helpers"
	portfolio "saham-app/model/portofolio"
	"saham-app/model/saham"
	"saham-app/model/transaction"
	"saham-app/model/user"
)

/*
 * Sell saham
 *
 */
func ProcessSell(namaPerusahaan string, lotJual int) {
	userID := user.UserLogin.ID

	sahamNow := saham.FindSahamByName(namaPerusahaan)
	if sahamNow == nil {
		helpers.GetMessages("❌ Saham tidak ditemukan.")
		return
	}

	port := portfolio.GetPortfolio(userID)
	data, exists := port[namaPerusahaan]
	if !exists || lotJual <= 0 || lotJual > data.TotalLot {
		helpers.GetMessages("❌ Jumlah lot tidak valid")
		return
	}

	totalJual := lotJual * 100 * sahamNow.Price_Per_Share
	user.UserLogin.Saldo += totalJual

	transaksiJual := transaction.Transaction{
		UserID:         userID,
		NamaPerusahaan: namaPerusahaan,
		JumlahLot:      lotJual,
		HargaPerLembar: sahamNow.Price_Per_Share,
		Total:          totalJual,
		Tipe:           "SELL",
	}
	transaction.AppendToHistory(transaksiJual)

	sisaLot := lotJual
	for i := range transaction.Transactions {
		t := &transaction.Transactions[i]
		if t.UserID == userID && t.NamaPerusahaan == namaPerusahaan {
			if t.JumlahLot <= sisaLot {
				sisaLot -= t.JumlahLot
				t.JumlahLot = 0
			} else {
				t.JumlahLot -= sisaLot
				break
			}
		}
	}

	var updated []transaction.Transaction
	for _, t := range transaction.Transactions {
		if t.JumlahLot > 0 {
			updated = append(updated, t)
		}
	}
	transaction.Transactions = updated

	user.SaveUsers()
	transaction.SaveTransactions()

	pesan := fmt.Sprintf("✅ Berhasil menjual %d lot saham %s seharga %d", lotJual, namaPerusahaan, totalJual)
	helpers.GetMessages(pesan)
}
