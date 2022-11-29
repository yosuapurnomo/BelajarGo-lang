package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func getHello(name string) string {
	return "Hello " + name
}

func getHelloBro(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello" + name
	}
}

func main() {
	sayHello()
	sayHelloTo("Yosua", "Purnomo")
	hello := getHello("Yosua")
	fmt.Println(hello)
	helloBro := getHelloBro("")
	fmt.Println(helloBro)
}
