package main

import (
	"Belajar-Golang/ImportPackage/helper"
	"fmt"
)

func main() {
	helper.SayHello("Yosua")
	fmt.Println(helper.GetDataBase())
	// Error karena function tersebut tidak bisa diakses diluar package
	// helper.sayHaii("Yosua")
}
