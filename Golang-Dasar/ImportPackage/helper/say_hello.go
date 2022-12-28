package helper

import "fmt"

var Nama = "Purnomo"

/** Jika Menggunakan awalan huruf kecil
maka variable tersebut tidak bisa digunakan diluar package*/
func sayHaii(name string) {
	fmt.Println("Haii,", name)
}

/**Menggunakan awalan huruf besar agar variabel/function dapat
diakses diluar pacakage*/
func SayHello(name string) {
	fmt.Println("Hello,", name)
}
