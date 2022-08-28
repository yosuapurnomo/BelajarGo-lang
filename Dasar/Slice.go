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
	makeSlice := make([]string, 2, 4)
	makeSlice[0] = "Yosua"
	makeSlice[1] = "Purnomo"
	// makeSlice[2] = "Yes"

	fmt.Println(makeSlice)
	makeSlice2 := append(makeSlice, "yes")
	// fmt.Println(makeSlice)
	fmt.Println(makeSlice2)
	fmt.Println(cap(makeSlice2))
	makeSlice3 := append(makeSlice2, "no")
	fmt.Println(makeSlice3)
	fmt.Println("Capasitas makeSlice3 =", cap(makeSlice3))
	makeSlice4 := append(makeSlice3, "Tambah", "lagi", "lagi")
	fmt.Println(makeSlice4)
	fmt.Println(cap(makeSlice4))
	fmt.Println(makeSlice)
	fmt.Println(cap(makeSlice))

	copySlice := make([]string, len(makeSlice3), cap(makeSlice3))
	copy(copySlice, makeSlice3)

	fmt.Println(copySlice)
}