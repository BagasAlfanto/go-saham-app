package controllers

import (
	"fmt"
	"saham-app/handlers/purchase"
	"saham-app/handlers/selling"
	"saham-app/helpers"
	"saham-app/model/saham"
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
			ShowSaham()
		case 2:
			SearchingSaham()
		case 3:
			saham.SortAscending()
		case 4:
			saham.SortDescending()
		case 5:
			isRunning = true
		}

		if choice != 5 {
			choice = 0
		}
		saham.UpdatePrice()
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
		fmt.Printf("| %-15s %-30s %-12d |\n", daftarSaham.StockCode, daftarSaham.CompanyName, daftarSaham.Price_Per_Share)
	}
	fmt.Println("===============================================================")
	saham.UpdatePrice()

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
	result := saham.Searching(data)
	helpers.DisplayShowSaham()
	fmt.Println(result)
	fmt.Println("===============================================================")

	helpers.ConfirmationScreen()
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
		fmt.Printf("| %-15s %-30s %-12d |\n", daftarSaham.StockCode, daftarSaham.CompanyName, daftarSaham.Price_Per_Share)
	}
	fmt.Println("===============================================================")

	fmt.Print("Masukan Kode atau nama perusahaan : ")
	fmt.Scan(&choices)

	selected := saham.FindSahamByCodeOrName(choices)
	if selected == nil {
		fmt.Println("❌ Saham tidak ditemukan.")
		helpers.ConfirmationScreen()
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

	userID := user.UserLogin.ID
	port := user.GetPortfolio(userID)

	if len(port) == 0 {
		helpers.GetMessages("❌ Kamu belum memiliki saham.")
		return
	}

	i := 1
	fmt.Println("===== Saham yang Dimiliki =====")
	for nama, data := range port {
		fmt.Printf("%d. %s - %d lot\n", i, nama, data.TotalLot)
		daftar = append(daftar, PortofolioItem{nama, data.TotalLot})
		i++
	}

	var pilihan int
	fmt.Print("Masukkan nomor saham yang ingin dijual: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > len(daftar) {
		helpers.GetMessages("❌ Pilihan tidak valid.")
		return
	}

	helpers.ClearScreen()

	namaPerusahaan := daftar[pilihan-1].Nama
	totalLot := daftar[pilihan-1].Lot

	fmt.Printf("Jumlah lot yang kamu miliki di %s : %d\n", namaPerusahaan, totalLot)
	fmt.Print("Masukkan jumlah lot yang ingin dijual: ")

	var lotJual int
	fmt.Scan(&lotJual)

	selling.ProcessSell(namaPerusahaan, lotJual)
}
