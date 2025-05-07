package saham

import (
	"fmt"
	"math/rand"
	"saham-app/helpers"
	"strings"
)

type Saham struct {
	IDSaham         int
	StockCode       string
	CompanyName     string
	Price_Per_Share int
}

var daftarSaham []Saham

var SearchSaham *Saham

func init() {
	daftarSaham = []Saham{
		{1, "AAPL", "Apple Inc.", 150},
		{2, "GOOGL", "Alphabet Inc.", 2800},
		{3, "AMZN", "Amazon.com Inc.", 3400},
		{4, "MSFT", "Microsoft Corp.", 299},
		{5, "TSLA", "Tesla Inc.", 700},
		{6, "FB", "Meta Platforms Inc.", 350},
	}
}

func GetSaham() []Saham {
	return daftarSaham
}

func ChangePricing() int {
	change := rand.Intn(201) - 100
	return change
}

func UpdatePrice() {
	random := rand.Intn(len(daftarSaham))
	priceChange := ChangePricing()

	daftarSaham[random].Price_Per_Share += priceChange
	if daftarSaham[random].Price_Per_Share < 0 {
		daftarSaham[random].Price_Per_Share = 0
	}
	if daftarSaham[random].Price_Per_Share > 5000 {
		daftarSaham[random].Price_Per_Share = 5000
	}
	daftarSaham[random].IDSaham = random + 1
}

func Searching(data string) string {
	for _, saham := range daftarSaham {
		if strings.EqualFold(saham.StockCode, data) || strings.Contains(strings.ToLower(saham.CompanyName), strings.ToLower(data)) {
			SearchSaham = &saham
			return fmt.Sprintf(
				"ID: %d\nKode Saham: %s\nPerusahaan: %s\nHarga per Lembar: %d",
				saham.IDSaham, saham.StockCode, saham.CompanyName, saham.Price_Per_Share,
			)
		}
	}
	return "Saham tidak ditemukan."
}

func ShortAscending() {
	var back string

	n := len(daftarSaham)
	for i := 1; i < n; i++ {
		key := daftarSaham[i]
		j := i - 1

		for j >= 0 && daftarSaham[j].Price_Per_Share > key.Price_Per_Share {
			daftarSaham[j+1] = daftarSaham[j]
			j--
		}
		daftarSaham[j+1] = key
	}
	
	helpers.DisplayShowSaham()
	for _, saham := range daftarSaham {
		fmt.Printf("| %-15s %-30s %-12d |\n", saham.StockCode, saham.CompanyName, saham.Price_Per_Share)
	}
	fmt.Println("===============================================================")

	for back != "ya" {
		fmt.Println("Kembali ke menu utama? (ya)")
		fmt.Scan(&back)
		UpdatePrice()
		helpers.ClearScreen()
	}
}

func ShortDescending() {
	var back string

	n := len(daftarSaham)
	for i := 1; i < n; i++ {
		key := daftarSaham[i]
		j := i - 1

		for j >= 0 && daftarSaham[j].Price_Per_Share < key.Price_Per_Share {
			daftarSaham[j+1] = daftarSaham[j]
			j--
		}
		daftarSaham[j+1] = key
	}

	helpers.DisplayShowSaham()
	for _, saham := range daftarSaham {
		fmt.Printf("| %-15s %-30s %-12d |\n", saham.StockCode, saham.CompanyName, saham.Price_Per_Share)
	}
	fmt.Println("===============================================================")

	for back != "ya" {
		fmt.Println("Kembali ke menu utama? (ya)")
		fmt.Scan(&back)
		UpdatePrice()
		helpers.ClearScreen()

	}
}
