package purchase

import (
	"fmt"
	"saham-app/helpers"
	"saham-app/model/saham"
	"saham-app/model/transaction"
	"saham-app/model/user"
)

/*
 * Purchase saham
 *
 */
func Purchase(selectedSaham *saham.Saham) {
	var jumlahLot int
	fmt.Printf("Masukkan jumlah lot yang ingin dibeli (1 lot = 100 lembar): ")
	n, err := fmt.Scan(&jumlahLot)

	if n < 1 {
		helpers.GetMessages("❌ Minimal beli lot adalah 1")
		return
	}
	if n != 1 || err != nil {
		helpers.GetMessages("❌ Input harus berupa angka")
		return
	}

	totalHarga := selectedSaham.Price_Per_Share * 100 * jumlahLot

	if user.UserLogin.Saldo < totalHarga {
		helpers.GetMessages("❌ Saldo tidak cukup untuk melakukan pembelian ini.")
		return
	}

	user.UserLogin.Saldo -= totalHarga

	transaksi := transaction.Transaction{
		UserID:         user.UserLogin.ID,
		NamaPerusahaan: selectedSaham.CompanyName,
		JumlahLot:      jumlahLot,
		HargaPerLembar: selectedSaham.Price_Per_Share,
		Total:          totalHarga,
		Tipe:           "BUY",
	}

	transaction.InsertTransaction(transaksi)
	transaction.AppendToHistory(transaksi)

	user.SaveUsers()
	helpers.GetMessages("✅ Pembelian saham berhasil")
}
