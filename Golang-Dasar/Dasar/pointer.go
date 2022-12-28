package main

import "fmt"

type DataDiri struct {
	NamaDepan, NamaBelakang, Alamat string
}

func main() {
	// Pemrograman Go menerapkan pass by value
	// jika sebuah variabel baru dibarikan nilai dari variabel lain maka variabel tersebut akan mengcopykan ke variabel baru tersebut
	yosua := DataDiri{"yosua", "purnomo", "jogoloyo"}
	// var yosua *DataDiri = &DataDiri{"yosua", "purnomo", "jogoloyo"}
	purnomo := yosua
	fmt.Println(yosua)
	fmt.Println("------------------------------------------------------------")
	// disaat variabel baru (purnomo) melakukan perubahan maka variabel lama (yosua) tidak ikut berubah
	purnomo.Alamat = "Wiyung"
	fmt.Println(yosua)
	fmt.Println(purnomo)

	fmt.Println("=========================================================")
	// fungsi pointer adalah untuk merubah pass by value menjadi pass by reference dengan menambahkan kode & didepan variabel
	bambang := &yosua
	fmt.Println(yosua)
	fmt.Println("------------------------------------------------------------")
	// ketika ada perubahan data maka variabel sebelumnya akan ikut terubah
	bambang.NamaDepan = "Bambang"
	fmt.Println(yosua)
	fmt.Println(purnomo)
	fmt.Println(bambang)

	// dengan memberikan kode * didepan variabel berfungsi untuk merubah data pada semua variabel yang memiliki reference dengan variabel tersebut
	fmt.Println("=========================================================")
	yuli := &yosua
	yuli.NamaDepan = "Yuli"
	fmt.Println(yuli)
	fmt.Println("------------------------------------------------------------")
	*bambang = DataDiri{"Bambang", "Tio", "Gunungsari"}
	fmt.Println(yosua)
	fmt.Println(bambang)
	fmt.Println(yuli)

	// jika ingin membuat pointer baru dengan tidak memiliki nilai
	joko := new(DataDiri)
	fmt.Println("=========================================================")
	fmt.Println(joko)
	joko.NamaDepan = "Joko"
	fmt.Println("------------------------------------------------------------")
	fmt.Println(joko)
}
