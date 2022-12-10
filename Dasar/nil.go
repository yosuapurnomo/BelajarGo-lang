package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var person map[string]string = NewMap("Yosua")

	if person == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println("Nama", person["name"])
	}
}
