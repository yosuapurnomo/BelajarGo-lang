package main

import "fmt"

func notif() {
	fmt.Println("Selesai memanggil function notif")
}

func runApplication(value int) {
	// fungsi defer adalah menjalankan sebuah proses lain disaat proses sebuah function berakhir meskipun dalam kondisi error
	// Contoh: function notif akan tetap dijalankan meskipun hasil pembagian result mengalami error
	defer notif()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println(result)
}

func main() {
	runApplication(2)
	fmt.Println("---------------------------------------------------------")
	runApplication(0)
}
