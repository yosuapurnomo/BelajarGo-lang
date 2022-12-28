package main

import "fmt"

func getName(name string) string {
	return "Ini Namammu " + name
}

func main() {
	sayName := getName

	result := getName("Purnomo")
	fmt.Println(result)
	fmt.Println(sayName("Yosua"))
}
