package main

import (
	"errors"
	"fmt"
)

func pembagi(nilai int, bagi int) (int, error) {
	if bagi == 0 {
		return 0, errors.New("Nilai Pembagi tidak boleh 0")
	} else {
		return nilai / bagi, nil
	}
}

func main() {
	var contohError error = errors.New("Apilasi Error")
	fmt.Println(contohError)

	hasil, err := pembagi(10, 2)
	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println(hasil, err)
	}
}
