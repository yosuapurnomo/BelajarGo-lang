package main

import "fmt"

func main() {
	name := "purnomo"

	if name == "yosua" {
		fmt.Println("Hai", name)
	} else if name == "Purnomo"{
		fmt.Println("Hai, Yosua", name)
	} else{
		fmt.Println("Nama bukan yosua purnomo melaikan", name)
	}

	// If Short Statement

	if lenght := len(name); lenght > 5 {
		fmt.Println("Panjang nama lebih dari 5 =", lenght)
	} else if lenght == 5{
		fmt.Println("Panjang nama sama dengan 5 =", lenght)
	} else{
		fmt.Println("Panjang nama kurang dari 5 = ", lenght)
	}
}