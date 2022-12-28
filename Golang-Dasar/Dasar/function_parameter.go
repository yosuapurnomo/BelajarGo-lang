package main

import "fmt"

type Filter func(string) string

func sayHelloType(nama string, filter Filter) string {
	return filter(nama)
}

func sayHello(nama string, filter func(string) string) string {
	return filter(nama)
}

func filterName(nama string) string {
	if nama == "anjing" {
		return "..."
	} else {
		return nama
	}
}

func main() {
	filter := sayHello("Yosua", filterName)
	filterLagi := sayHello("anjing", filterName)
	filterType := sayHelloType("Purnomo", filterName)

	fmt.Println("Nama ", filter)
	fmt.Println("Nama ", filterLagi)
	fmt.Println("Nama ", filterType)
}
