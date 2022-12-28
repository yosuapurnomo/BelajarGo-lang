package main

import "fmt"

type Block func(string) bool

func validationUser(nama string, block Block) string {
	if block(nama) {
		return "Welcome " + nama
	} else {
		return "You are Blocked " + nama
	}
}

func main() {
	blacklist := func(nama string) bool {
		return nama == "admin"
	}

	fmt.Println(validationUser("admin", blacklist))
	fmt.Println(validationUser("yosua", blacklist))

	fmt.Println(validationUser("admin", func(nama string) bool {
		return nama == "admin"
	}))
	fmt.Println(validationUser("Purnomo", func(nama string) bool {
		return nama == "admin"
	}))
}
