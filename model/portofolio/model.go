package portfolio

import (
	"saham-app/model/transaction"
)

/*
 * Mendapatkan data saham dari user
 *
 */
func GetPortfolio(userID int) map[string]struct {
	TotalLot   int
	TotalModal int
} {
	result := make(map[string]struct {
		TotalLot   int
		TotalModal int
	})

	for _, t := range transaction.Transactions {
		if t.UserID == userID {
			val := result[t.NamaPerusahaan]
			val.TotalLot += t.JumlahLot
			val.TotalModal += t.Total
			result[t.NamaPerusahaan] = val
		}
	}

	return result
}
