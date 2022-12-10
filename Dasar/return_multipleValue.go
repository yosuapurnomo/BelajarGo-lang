package main

import "fmt"

func getFullName() (string, string) {
	return "yosua", "purnomo"
}

func getName() (string, string) {
	return "Yosua", "Purnomo"
}

func getNameComplete() (namaAwal string, namaTengah string, namaAkhir string) {
	namaAwal = "YOSUA"
	namaAkhir = "PURNOMO"
	namaTengah = "WOWWWW"

	return
}

func main() {
	first, last := getFullName()

	fmt.Println(first, last)

	firstName, _ := getName()

	_, lastName := getName()

	fmt.Println(firstName, lastName)

	firstName, middelName, lastName := getNameComplete()

	fmt.Println(firstName, middelName, lastName)
}
