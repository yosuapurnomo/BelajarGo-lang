package main

import "fmt"

/**
Struct adalah template data atau prototype data yang artinya tidak bisa langsung diisikan data
namun bisa membuat data/object dari struct yang telah dibuat
*/

type Biodata struct {
	Nama, Alamat string
	Umur         int
}

func (biodata Biodata) sayHi(name string) {
	fmt.Println("Hi", name, ", my name is", biodata.Nama)
}

func main() {
	// Deklarasi variable untuk struct cara 1
	var biodata Biodata
	// Pemberian value pada struct yang telah dibuat
	biodata.Nama = "Yosua Purnomo"
	biodata.Alamat = "Jogoloyo"
	biodata.Umur = 23

	fmt.Println(biodata)
	fmt.Println("------------------")
	fmt.Println(biodata.Nama)
	fmt.Println(biodata.Alamat)
	fmt.Println(biodata.Umur)
	fmt.Println("------------------")

	// Cara 2
	toni := Biodata{
		Nama:   "Toni",
		Alamat: "Wiyung",
		Umur:   25,
	}
	fmt.Println(toni)
	fmt.Println("------------------")
	// Cara 3 (Sangat tidak disarankan karena jika terjadi perubahan pada struct maka akan terjadi error jika deklarasi tidak dirubah disesuaikan)
	// Parameter sesuaikan dengan urutan
	wawan := Biodata{"Wawan", "Gunungsari", 30}
	fmt.Println(wawan)
	fmt.Println("------------------")
	// Struct Function / Struct Method
	biodata.sayHi("Joko")
}
