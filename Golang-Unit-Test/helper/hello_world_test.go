package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := SayHello("Yosua")

	if result != "Hello Yosua" {
		panic("Error, is not say Hello Yosua")
	}
}

// Menggagalkan tes sebenarnya tidak disarankan menggunakan panic, namun cara terbaik adalah menggunakan
// function Fail(), FailNow(), error,

// Fail()

func TestFail(t *testing.T) {
	result := SayHello("Yosua")

	if result != "Hai Yosua" {
		t.Fail()
	}
	fmt.Println("Fail Done")
}

func TestOtherFailNow(t *testing.T) {
	result := SayHello("Yosua")

	if result != "Hai Yosua" {
		t.FailNow()
	}
	fmt.Println("FailNow Done")
}
