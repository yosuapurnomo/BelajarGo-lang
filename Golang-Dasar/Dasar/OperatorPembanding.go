package main

import "fmt"

func main() {
	var name1 = "yosua"
	var name2 = "yosua"

	var checkName bool = name1 == name2
	fmt.Println(checkName)

	var name3 = "purnomo"

	var checkLastName bool = name2 == name3
	
	fmt.Println(checkLastName)
	fmt.Println(name1 != name3)

	var nilai1 = 100
	var nilai2 = 50

	var checkNilai = nilai1 > nilai2

	fmt. Println(checkNilai)
	fmt. Println(nilai1 < nilai2)
}