package main

import "fmt"

func main(){
	dataMap := map[int]string{
		1:"satu",
		2:"dua",
	}

	fmt.Println(dataMap)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Yosua Purnomo"
	book["salah"] = "delete"
	fmt.Println(book)

	delete(book, "salah")

	fmt.Println(book)
}