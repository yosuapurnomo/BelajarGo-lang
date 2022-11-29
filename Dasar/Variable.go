package main

import "fmt"

func main(){
	var lastName string
	var firstName = "Yosua"
	lastName = "Purnomo"
	fmt.Println(firstName)
	fmt.Println(lastName)

	// deklarasi secara langsung tanpa memberikan tipe data dan mengikutin sesuai dengan tipe data value yang diberikan
	age := 23
	fmt.Println("Umur", age)

	// deklarasi banyak variable
	var (
		country = "Indonesia"
		city = "Surabaya"
	)

	fmt.Println("Country", country)
	fmt.Println("City", city)

	const (
		nik = 12389213893
	)
	fmt.Println(nik)
}