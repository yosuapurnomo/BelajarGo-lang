package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(run bool) {
	defer endApp()
	if run {
		// Panic function berfungsi untuk menghentikan aplikasi yang berjalan dengan mengembalikan sebuah error
		// dan terdapat defer function yang tetap dijalankan diakhir proses
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi Jalan")
}

func main() {
	runApp(false)
	fmt.Println("------------------------------------------------")
	runApp(true)
}
