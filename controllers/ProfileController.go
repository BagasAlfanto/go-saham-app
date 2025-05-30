package controllers

import (
	"fmt"
	"saham-app/helpers"
	"saham-app/model/saham"
	"saham-app/model/transaction"
	"saham-app/model/user"
)

/*
 * Menampilkan halaman profile user
 *
 */
func ShowProfile() {
	var choice int

	for isRunning := false; !isRunning; {
		helpers.ClearScreen()

		profile := user.GetUser()
		fmt.Println("===== Profil Pengguna =====")
		fmt.Println("Username :", profile.Username)
		fmt.Println("Saldo    :", helpers.NominalFormat(profile.Saldo))
		fmt.Println()

		for choice < 1 || choice > 3 {
			fmt.Println("1. Lihat Riwayat Transaksi")
			fmt.Println("2. Lihat Portofolio")
			fmt.Println("3. Kembali ke menu sebelumnya")

			fmt.Print("Pilih opsi: ")
			fmt.Scan(&choice)
		}

		switch choice {
		case 1:
			ShowTransactionHistory()
		case 2:
			ShowPortfolio()
		case 3:
			isRunning = true
		}

		if choice != 3 {
			choice = 0
		}
		saham.ChangePricing()
	}
}

/*
 * Menampilkan transaksi user
 *
 */
func ShowTransactionHistory() {
	helpers.ClearScreen()

	profile := user.GetUser()
	fmt.Println("===== Riwayat Transaksi Saham =====")
	fmt.Println("Username :", profile.Username)
	fmt.Println("Saldo    :", helpers.NominalFormat(profile.Saldo))
	fmt.Println()

	transaksiUser := transaction.GetHistoryByUserID(user.UserLogin.ID)
	if len(transaksiUser) == 0 {
		fmt.Println("❌ Belum ada transaksi saham.")
	} else {
		fmt.Printf("%-5s %-10s %-30s %-10s %-15s %-15s\n", "No", "Tipe", "Perusahaan", "Lot", "Harga/Lembar", "Total")
		fmt.Println("------------------------------------------------------------------------------------------")

		for i, t := range transaksiUser {
			i++
			fmt.Printf("%-5d %-10s %-30s %-10d %-15s %-15s\n",
				i, t.Tipe, t.NamaPerusahaan, t.JumlahLot, helpers.NominalFormat(t.HargaPerLembar), helpers.NominalFormat(t.Total))
		}
	}

	helpers.ConfirmationScreen()
}

/*
 * Menampilkan portofolio user
 *
 */
func ShowPortfolio() {
	helpers.ClearScreen()
	saham.UpdatePrice()

	profile := user.GetUser()
	fmt.Println("===== Portofolio Pengguna =====")
	fmt.Println("Username :", profile.Username)
	fmt.Println("Saldo    :", helpers.NominalFormat(profile.Saldo))
	fmt.Println()

	transaksiUser := transaction.GetTransactionsByUserID(user.UserLogin.ID)
	if len(transaksiUser) == 0 {
		fmt.Println("❌ Tidak ada saham di portofolio.")
		helpers.ConfirmationScreen()
		return
	}

	portofolio := make(map[string]int)
	modalMap := make(map[string]int)

	for _, t := range transaksiUser {
		portofolio[t.NamaPerusahaan] += t.JumlahLot
		modalMap[t.NamaPerusahaan] += t.Total
	}

	fmt.Printf("%-5s %-30s %-10s %-15s %-15s %-15s\n", "No", "Perusahaan", "Lot", "Harga Saat Ini", "Nilai", "Untung/Rugi")
	fmt.Println("-----------------------------------------------------------------------------------------------------")

	i := 1
	for nama, totalLot := range portofolio {

		s := saham.FindSahamByName(nama)
		if s == nil {
			fmt.Printf("%-30s %-10d %-15s %-15s %-15s\n", nama, totalLot, "-", "-", "-")
			continue
		}

		hargaSekarang := s.Price_Per_Share
		nilai := hargaSekarang * 100 * totalLot
		modal := modalMap[nama]
		selisih, status := saham.CalculateDifference(nilai, modal)

		fmt.Printf("%-5d %-30s %-10d %-15s %-15s %s: %s\n", i, nama, totalLot, helpers.NominalFormat(hargaSekarang), helpers.NominalFormat(nilai), status, helpers.NominalFormat(selisih))
		i++

	}

	helpers.ConfirmationScreen()
}
