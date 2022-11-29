package main

import "fmt"

func main() {
	name := "selly"

	switch name {
	case "yosua":
		fmt.Println("ini Yosua =", name)
	case "purnomo":
		fmt.Println("ini Purnomo =", name)
	default: fmt.Println("Default, nama =", name)
	}
	
	// Switch Short Statement

	switch lenght := len(name); lenght > 5{
	case true: fmt.Println("Nama lebih dari 5 =", lenght)
	case false: fmt.Println("Nama kurang dari sama dengan 5 =", lenght)
	}

	// Switch no Condition

	lenght := len(name)
	switch{
	case lenght > 5: fmt.Println("Nama lebih dari 5 =", lenght)
	case lenght < 5: fmt.Println("Nama kurang dari 5 =", lenght)
	default : fmt.Println("Nama = 5 =", lenght)
	}
}