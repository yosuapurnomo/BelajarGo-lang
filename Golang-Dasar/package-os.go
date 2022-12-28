package main

import (
	"fmt"
	"os"
)

func main() {
	// Run application: go run package-os.go __argumen__

	args := os.Args
	fmt.Println(args)

	fmt.Println(args[1])
	fmt.Println(args[2])


}
