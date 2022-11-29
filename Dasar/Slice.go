package main

import "fmt"

func main() {
	var data = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"Nopember",
		"Desember",
	}

	fmt.Println(data)

	slice := data[2:8]
	fmt.Println(slice)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// saat data pada array terubah, maka data slice pun juga ikut berubah
	data[4] = "Dirubah"
	fmt.Println(slice)

	// saat data pada slice terubah, maka data array pun juga ikut berubah
	slice[0] = "Slice terubah"
	fmt.Println(data)

	slice2 := data[5:]
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slice3 := append(slice2, "Yosua")
	fmt.Println(slice3)
	slice3[1] = "Bukan bulan"

	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(data)
	fmt.Println("-----------------------------------------------------")

	// Make slice dari awal
	// Kapasitas Slice dapat diatur dari awal pendeklarasian data dengan value kosong
	makeSlice := make([]string, 3, 3)
	fmt.Println("Capasitas Data Slice Awal =", cap(makeSlice))
	fmt.Println("Data Awal kosong =", makeSlice)
	makeSlice[0] = "1"
	makeSlice[1] = "2"
	makeSlice[2] = "3"

	fmt.Println("Slice awal diberikan 3 value =", makeSlice)
	makeSlice = append(makeSlice, "4")
	// fmt.Println(makeSlice)
	fmt.Println("Ini Datanya setelah ditambahkan", makeSlice)
	fmt.Println("Ini Capasitasnya", cap(makeSlice))
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	// Disaat capasitas data yang ditampung melebihi capasitasnya, maka akan dibuatkan data array baru yang nantinya
	// jika ada perubahan pada array baru tersebut tidak akan merubah data array sebelumnya
	makeSlice3 := append(makeSlice, "5")
	fmt.Println("Ini Data Slice3 =", makeSlice3)
	fmt.Println("Capasitas makeSlice3 =", cap(makeSlice3))
	makeSlice3 = append(makeSlice3, "6", "7", "8")
	fmt.Println("Ini Data Slice4 =", makeSlice3)
	fmt.Println("Capasitas makeSlice4 =", cap(makeSlice3))
	fmt.Println("Ini Data Slice awal =", makeSlice)
	fmt.Println("Capasitas makeSlice awal =", cap(makeSlice))

	copySlice := make([]string, len(makeSlice3), cap(makeSlice3))
	// copy(variable baru/variable uang akan menerima, variable lama/variable yang akan memberikan value)
	copy(copySlice, makeSlice3)

	fmt.Println("Hasil Copy data slice =", copySlice)
}
