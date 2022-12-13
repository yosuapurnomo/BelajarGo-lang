package main

import "fmt"

type Biodata struct {
	NamaDepan, NamaBelakang, Alamat string
}

func UbahNamaDepan(bio *Biodata) {
	bio.NamaDepan = "Dodik"
}

func UbahNamaBelakang(bio Biodata) {
	bio.NamaBelakang = "didi"
}

func main() {
	var dodik Biodata = Biodata{
		NamaDepan:    "",
		NamaBelakang: "Purnomo",
		Alamat:       "Jogoloyo",
	}

	// Tanpa Pointer.. Data struct tidak terubah melalui function tanpa menambahkan return data pada function
	UbahNamaBelakang(dodik)
	fmt.Println(dodik)

	// Menggunakan Pointer.. Data Struct dapat dirubah melalui function tanpa menambahkan return data pada function
	UbahNamaDepan(&dodik)
	fmt.Println(dodik)

}
