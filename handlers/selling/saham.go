package selling

import (
	"fmt"
	"saham-app/helpers"
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

	// Validasi saham
	sahamNow := saham.FindSahamByName(namaPerusahaan)
	if sahamNow == nil {
		helpers.GetMessages("❌ Saham tidak ditemukan.")
		return
	}

	// Ambil total lot dimiliki
	port := user.GetPortfolio(userID)
	data, exists := port[namaPerusahaan]
	if !exists || lotJual <= 0 || lotJual > data.TotalLot {
		helpers.GetMessages("❌ Jumlah lot tidak valid atau tidak punya saham tersebut.")
		return
	}

	// Hitung nilai jual
	totalJual := lotJual * 100 * sahamNow.Price_Per_Share
	user.UserLogin.Saldo += totalJual

	// Catat transaksi SELL ke history
	transaksiJual := transaction.Transaction{
		UserID:         userID,
		NamaPerusahaan: namaPerusahaan,
		JumlahLot:      lotJual,
		HargaPerLembar: sahamNow.Price_Per_Share,
		Total:          totalJual,
		Tipe:           "SELL",
	}
	transaction.AppendToHistory(transaksiJual)

	// Kurangi dari portofolio (hapus dari Transactions)
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

	// Filter ulang Transactions
	var updated []transaction.Transaction
	for _, t := range transaction.Transactions {
		if t.JumlahLot > 0 {
			updated = append(updated, t)
		}
	}
	transaction.Transactions = updated

  // Simpan ke database dan history transaksi
	user.SaveUsers()
	transaction.SaveTransactions()

	pesan := fmt.Sprintf("✅ Berhasil menjual %d lot saham %s seharga %d", lotJual, namaPerusahaan, totalJual)
	helpers.GetMessages(pesan)
}
