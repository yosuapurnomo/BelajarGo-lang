package main

import "fmt"

type Orang interface {
	GetName() string
	GetAlamat() string
}

type Biodata struct {
	Nama   string
	Alamat string
}

func (biodata Biodata) GetName() string {
	return biodata.Nama
}

func (biodata Biodata) GetAlamat() string {
	return biodata.Alamat
}

type OrangLain struct {
	Nama   string
	Alamat string
}

func (biodata OrangLain) GetName() string {
	return biodata.Nama
}

func (biodata OrangLain) GetAlamat() string {
	return biodata.Alamat
}

func sayHello(orang Orang) {
	fmt.Println("Hello", orang.GetName(), "rumahmu di", orang.GetAlamat())
}

func main() {
	// Cara 1
	var yosu Biodata
	yosu.Nama = "Yosua"
	yosu.Alamat = "jogoloyo"
	sayHello(yosu)

	// Cara 2
	joko := OrangLain{
		Nama:   "Joko",
		Alamat: "Wiyung",
	}
	sayHello(joko)

}
