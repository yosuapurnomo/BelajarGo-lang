package main

import "fmt"

type Man struct {
	NamaDepan, NamaBelakang string
}

// Cara menjadikan struct function sebagai pointer
func (man *Man) UbahNama() {
	man.NamaDepan = "Yosua"
	man.NamaBelakang = "Purnomo"
}

func main() {
	yosua := Man{
		NamaDepan: "Yosua",
	}

	yosua.UbahNama()

	fmt.Println(yosua)
}
