package main

import "fmt"

func TypeAcak() interface{} {
	return true
}

func main() {
	var result interface{} = TypeAcak()
	// Type Assertion berfungsi untuk mengubah secara paksa pengembalian pada tipe data pada sebuah interface kosong
	// Jika pengembalian tidak sesuai dengan tipe data maka akan terjadi panic
	// fmt.Println(result.(string))

	// Maka solusinya adalah menggunakan switch case

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "Is String")
	case int:
		fmt.Println("Value", value, "Is Integer")
	default:
		fmt.Println("Value", value, "Is Unknow")
	}
}
