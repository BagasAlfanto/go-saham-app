package controllers

import (
	"fmt"
	"saham-app/handlers/purchase"
	"saham-app/handlers/selling"
	"saham-app/helpers"
	"saham-app/model/saham"
	"saham-app/model/transaction"
	"saham-app/model/user"
)

type PortofolioItem struct {
	Nama string
	Lot  int
}

var daftar []PortofolioItem

/*
 * Menampilkan halaman menu saham
 *
 */
func SahamMenu() {
	var choice int

	for isRunning := false; !isRunning; {
		helpers.ClearScreen()
		helpers.DisplaySaham()

		for choice < 1 || choice > 5 {
			fmt.Print("Masukan Pilihan Anda : ")
			fmt.Scan(&choice)
		}

		switch choice {
		case 1:
			saham.UpdatePrice()
			ShowSaham()
		case 2:
			SearchingSaham()
		case 3:
			saham.UpdatePrice()
			saham.SortAscending()
		case 4:
			saham.UpdatePrice()
			saham.SortDescending()
		case 5:
			isRunning = true
		}

		if choice != 5 {
			choice = 0
		}
	}

}

/*
 * Menampilkan data semua saham
 *
 */
func ShowSaham() {
	helpers.ClearScreen()
	helpers.DisplayShowSaham()

	for _, daftarSaham := range saham.GetSaham() {
		fmt.Printf("| %-15s %-30s %-12s |\n", daftarSaham.SahamCode, daftarSaham.CompanyName, helpers.NominalFormat(daftarSaham.Price_Per_Share))
	}
	fmt.Println("===============================================================")

	helpers.ConfirmationScreen()
}

/*
 * Menampilkan saham yang dicari
 *
 */
func SearchingSaham() {
	helpers.ClearScreen()
	var data string
	fmt.Println("================== Cari Saham ==================")
	fmt.Print("Masukan Kode atau Nama Perusahaan : ")
	fmt.Scan(&data)

	helpers.ClearScreen()
	ada, result := saham.Searching(data)
	if ada {
		helpers.DisplayShowSaham()
		fmt.Print(result)
		fmt.Println("===============================================================")
		helpers.ConfirmationScreen()
	} else {
		helpers.GetMessages(result)
	}

}

/*
 * Menampilkan proses beli saham
 *
 */
func BuyingSaham() {
	var choices string
	helpers.ClearScreen()
	helpers.DisplayShowSaham()

	for _, daftarSaham := range saham.GetSaham() {
		fmt.Printf("| %-15s %-30s %-12s |\n", daftarSaham.SahamCode, daftarSaham.CompanyName, helpers.NominalFormat(daftarSaham.Price_Per_Share))
	}
	fmt.Println("===============================================================")

	fmt.Print("Masukan Kode atau nama perusahaan : ")
	fmt.Scan(&choices)

	selected := saham.FindSahamByCodeOrName(choices)
	if selected == nil {
		helpers.GetMessages("❌ Saham tidak ditemukan.")
		return
	}

	purchase.Purchase(selected)
}

/*
 * Menampilkan proses jual saham
 *
 */
func SellSaham() {
	helpers.ClearScreen()
	daftar = nil

	userID := user.UserLogin.ID
	port := user.GetPortfolio(userID)

	if len(port) == 0 {
		helpers.GetMessages("❌ Kamu belum memiliki saham.")
		return
	}

	i := 1
	fmt.Println("===== Saham yang Dimiliki =====")
	for nama, data := range port {
		harga := 0
		for _, s := range saham.GetSaham() {
			if s.CompanyName == nama {
				harga = s.Price_Per_Share
				break
			}
		}
		fmt.Printf("%d. %s - %d lot (Harga per lembar: %s)\n", i, nama, data.TotalLot, helpers.NominalFormat(harga))
		daftar = append(daftar, PortofolioItem{nama, data.TotalLot})
		i++
	}

	var pilihan int
	fmt.Print("Masukkan nomor saham yang ingin dijual: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > len(daftar) {
		helpers.GetMessages("❌ Saham tidak ditemukan.")
		return
	}

	helpers.ClearScreen()

	namaPerusahaan := daftar[pilihan-1].Nama
	totalLot := 0
	for _, t := range transaction.Transactions {
		if t.UserID == user.UserLogin.ID && t.NamaPerusahaan == namaPerusahaan {
			totalLot += t.JumlahLot
		}
	}

	fmt.Printf("Jumlah lot yang kamu miliki di %s : %d\n", namaPerusahaan, totalLot)
	fmt.Print("Masukkan jumlah lot yang ingin dijual: ")

	var lotJual int
	fmt.Scan(&lotJual)

	selling.ProcessSell(namaPerusahaan, lotJual)
}
