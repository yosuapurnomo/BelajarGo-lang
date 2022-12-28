package main

import "fmt"

func hitungFaktorial(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * hitungFaktorial(number-1)
	}
}

func main() {
	faktor := hitungFaktorial(10)
	fmt.Println(faktor)
}
