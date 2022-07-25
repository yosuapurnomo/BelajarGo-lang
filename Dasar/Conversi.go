package main

import "fmt"

func main(){
	name := "Yosua Purnomo"
	y := string(name[0])
	
	fmt.Println(name)
	fmt.Println(y)

	var umur128 = 128
	var umur int8 = int8(umur128)

	fmt.Println(umur128)
	fmt.Println(umur)
}