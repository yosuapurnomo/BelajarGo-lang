package main

import "fmt"

func endApp() {
	/** recover function berfungsi untuk menangkap error yang dihasilkan dari panic function agar
	aplikasi tidak langsung berhenti dan dapat ditangani
	recover akan mengembalikan nilai nill jika tidak terjadi panic */
	message := recover()
	if message != nil {
		fmt.Println("Message error:", message)
	}
	fmt.Println("Aplikasi Selesai")
}

func runApp(run bool) {
	defer endApp()
	if run {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi berjalan")
	// recover function harus ditempatkan pada defer function agar tetap dapat dijalankan disaat error
}

func main() {
	runApp(false)
	fmt.Println("-------------------------------------------")
	runApp(true)
}
