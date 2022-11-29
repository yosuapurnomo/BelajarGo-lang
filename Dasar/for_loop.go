package main

import "fmt"

func main() {
	counter := 1

	fmt.Println(1)
	for counter <= 5 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	fmt.Println(2)
	for counter := 1; counter <= 5; counter++{
		fmt.Println("Perulangan ke", counter)
	}

	fmt.Println("\nSlice\n")
	slice := []string{"Yosua", "Purnomo"}

	for i := 0; i < len(slice); i++{
		fmt.Println(slice[i])
	}

	for index, name := range slice{
		fmt.Println("index", index, "name", name)
	}
}