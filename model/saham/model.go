package saham

type Saham struct {
	NamaSaham string
	Harga     int
}

var daftarSaham []Saham

func init() {
	daftarSaham = []Saham{
		{"BRI", 1000},
		{"BCA", 2000},
		{"BNI", 3000},
	}
}

func GetSaham() []Saham {
	return daftarSaham
}
