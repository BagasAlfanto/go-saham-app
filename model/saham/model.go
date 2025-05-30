package saham

import (
	"fmt"
	"math/rand"
	"saham-app/helpers"
	"strings"
)

/*
 * Struct saham
 *
 */
type Saham struct {
	IDSaham         int
	SahamCode       string
	CompanyName     string
	Price_Per_Share int
}

var daftarSaham []Saham

var SearchSaham *Saham

/*
 * Load daftar saham dari database
 *
 */
func init() {
	if helpers.FileExists("daftarsaham.json") {
		content, err := helpers.ReadFile("daftarsaham.json")
		if err != nil {
			panic(err)
		}

		helpers.LoadFromJSON(content, &daftarSaham)
	} else {
		daftarSaham = []Saham{
			{1, "AAPL", "Apple Inc.", 150},
			{2, "GOOGL", "Google Inc.", 999},
			{3, "AMZN", "Amazon.com Inc.", 890},
			{4, "MSFT", "Microsoft Corp.", 299},
			{5, "TSLA", "Tesla Inc.", 700},
			{6, "FB", "Meta Platforms Inc.", 350},
		}
		content, err := helpers.SaveToJSON(daftarSaham)

		if err != nil {
			panic(err)
		}

		err = helpers.UpdateFile("daftarsaham.json", content)

		if err != nil {
			panic(err)
		}
	}

}

/*
 * Mendapatkan data saham
 *
 */
func GetSaham() []Saham {
	return daftarSaham
}

/*
 * Update harga saham
 *
 */
func ChangePricing() int {
	change := rand.Intn(201) - 100
	return change
}

/*
 * Implementasi update harga saham
 *
 */
func UpdatePrice() {
	for i := range daftarSaham {
		priceChange := ChangePricing()
		daftarSaham[i].Price_Per_Share += priceChange

		if daftarSaham[i].Price_Per_Share < 0 {
			daftarSaham[i].Price_Per_Share = 0
		}
		if daftarSaham[i].Price_Per_Share > 5000 {
			daftarSaham[i].Price_Per_Share = 5000
		}
	}
}

/*
 * Searching Saham
 *
 */
func Searching(data string) (bool, string) {
	helpers.ClearScreen()
	var result string
	found := false
	n := len(daftarSaham)

	// Sequential/Linear Search implementation
	for i := 0; i < n; i++ {
		// Case-insensitive comparison for both stock code and company name
		if strings.EqualFold(daftarSaham[i].SahamCode, data) || strings.Contains(strings.ToLower(daftarSaham[i].CompanyName), strings.ToLower(data)) {
			found = true
			result += fmt.Sprintf(
				"| %-15s %-30s %-12s |\n",
				daftarSaham[i].SahamCode,
				daftarSaham[i].CompanyName,
				helpers.NominalFormat(daftarSaham[i].Price_Per_Share),
			)
		}
	}

	if !found {
		return false, "❌ Saham tidak ditemukan."
	}

	return true, result
}

/*
 * Mengurutkan daftar saham dari harga terendah menggunakan insertion sort
 *
 */
func SortAscending() {
	sorted := make([]Saham, len(daftarSaham))
	copy(sorted, daftarSaham)

	helpers.ClearScreen()
	n := len(sorted)
	for i := 1; i < n; i++ {
		sort := sorted[i]
		j := i - 1

		for j >= 0 && sorted[j].Price_Per_Share > sort.Price_Per_Share {
			sorted[j+1] = sorted[j]
			j--
		}
		sorted[j+1] = sort
	}

	helpers.DisplayShowSaham()
	for _, saham := range sorted {
		fmt.Printf("| %-15s %-30s %-12s |\n", saham.SahamCode, saham.CompanyName, helpers.NominalFormat(saham.Price_Per_Share))
	}
	fmt.Println("===============================================================")

	helpers.ConfirmationScreen()
}

/*
 * Mengurutkan daftar saham dari harga tertinggi menggunakan selection sort
 *
 */
func SortDescending() {
	sorted := make([]Saham, len(daftarSaham))
	copy(sorted, daftarSaham)

	helpers.ClearScreen()
	n := len(sorted)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if sorted[j].Price_Per_Share > sorted[maxIdx].Price_Per_Share {
				maxIdx = j
			}
		}
		if maxIdx != i {
			sorted[i], sorted[maxIdx] = sorted[maxIdx], sorted[i]
		}
	}

	helpers.DisplayShowSaham()
	for _, saham := range sorted {
		fmt.Printf("| %-15s %-30s %-12s |\n", saham.SahamCode, saham.CompanyName, helpers.NominalFormat(saham.Price_Per_Share))
	}
	fmt.Println("===============================================================")

	helpers.ConfirmationScreen()
}

/*
 * Mencari saham by kode / nama perusahaan
 *
 */
func FindSahamByCodeOrName(input string) *Saham {
	for _, s := range daftarSaham {
		if strings.EqualFold(s.SahamCode, input) || strings.Contains(strings.ToLower(s.CompanyName), strings.ToLower(input)) {
			return &s
		}
	}
	return nil
}

/*
 * Mencari saham by nama perusahaan
 *
 */
func FindSahamByName(name string) *Saham {
	for _, s := range daftarSaham {
		if strings.EqualFold(s.CompanyName, name) {
			return &s
		}
	}
	return nil
}

/*
 * Menghitung selisih antara nilai sekarang dan modal
 *
 *
 */
func CalculateDifference(nilaiSekarang, modal int) (selisih int, status string) {
	selisih = nilaiSekarang - modal
	if selisih < 0 {
		status = "Rugi"
	} else {
		status = "Untung"
	}
	return
}

/*
 * Menyimpan perubahan harga saham
 *
 */
func SaveSaham() {
	content, err := helpers.SaveToJSON(daftarSaham)
	if err != nil {
		fmt.Println("❌ Gagal menyimpan saham:", err)
		return
	}

	err = helpers.UpdateFile("daftarsaham.json", content)
	if err != nil {
		fmt.Println("❌ Gagal update file daftarsaham.json:", err)
	}
}
