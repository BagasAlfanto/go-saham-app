package controllers

import (
	"fmt"
	"saham-app/helpers"
	"saham-app/model/saham"
	"saham-app/model/user"
)

func ShowProfile() {
	var back string
	helpers.ClearScreen()
	fmt.Println("===============================================================")
	fmt.Println("------------------------ PROFILE ------------------------------")
	fmt.Println("===============================================================")
	fmt.Println("Username		: ", user.GetUser().Username)
	fmt.Println("Saldo			: ", user.GetUser().Saldo)
	fmt.Println("Saham yang dimiliki	: ", user.GetUser().Saham)
	fmt.Println("===============================================================")

	for back != "ya" {
		fmt.Println("Kembali ke menu utama? (ya)")
		fmt.Scan(&back)
		saham.UpdatePrice()
		helpers.ClearScreen()
	}
}
