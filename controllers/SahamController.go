package controllers

import (
	"fmt"
	"saham-app/helpers"
	"saham-app/model/saham"
)

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
			saham.ShortAscending()
		case 4:
			saham.ShortDescending()
		case 5:
			isRunning = true
		}

		if choice != 5 {
			choice = 0
		}
	}

}

func ShowSaham() {
	var back string

	helpers.ClearScreen()
	helpers.DisplayShowSaham()
	
	for _, daftarSaham := range saham.GetSaham() {
		fmt.Printf("| %-15s %-30s %-12d |\n", daftarSaham.StockCode, daftarSaham.CompanyName, daftarSaham.Price_Per_Share)
	}
	fmt.Println("===============================================================")
	for back != "ya" {
		fmt.Println("Kembali ke menu utama? (ya)")
		fmt.Scan(&back)
		saham.UpdatePrice()
		helpers.ClearScreen()
	}
}

func SearchingSaham() {
	var data, back string
	fmt.Println("================== Cari Saham ==================")
	fmt.Print("Masukan Kode atau Nama Perusahaan : ")
	fmt.Scan(&data)

	result := saham.Searching(data)
	fmt.Println(result)

	for back != "ya" {
		fmt.Println("Kembali ke menu utama? (ya)")
		fmt.Scan(&back)
		saham.UpdatePrice()
		helpers.ClearScreen()
	}
}
