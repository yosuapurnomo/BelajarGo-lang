package main

import "fmt"

func main() {
	var name [2]string

	name[0] = "yosua"
	name[1] = "Purnomo"

	fmt.Println(name[1])
	fmt.Println(name[0])

	var name2 = [2]string{
		"Yosua", "Purnomo",
	}

	fmt.Println(name2)

	fmt.Println("Hitung Panjang Array", len(name2))

	var lagi [10]int

	fmt.Println("Hitung panjang array tanpa ada data/kosong", len(lagi))
}