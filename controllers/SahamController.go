package controllers

import (
	"fmt"
	"saham-app/helpers"
	"saham-app/model/saham"
)

func ShowSaham() {
	var back string
	helpers.ClearScreen()
	fmt.Println("================ Daftar Saham ===============")
	for _, daftarSaham := range saham.GetSaham() {
		fmt.Printf("Nama Saham : %s\n", daftarSaham.NamaSaham)
		fmt.Printf("Harga      : %d\n", daftarSaham.Harga)
		fmt.Println("----------------------------------------------")
	}

	for back != "ya" {
		fmt.Println("Kembali ke menu utama? (ya)")
		fmt.Scan(&back)
		helpers.ClearScreen()
	}
}
