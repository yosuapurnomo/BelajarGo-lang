package main

import "fmt"

func main(){
	a := 10
	b := 2+a
	fmt.Println(b)

	c := 5
	b +=c
	fmt.Println(b)
	fmt.Println(c)

	d := a*b
	fmt.Println(d)

	e := b/a
	fmt.Println(e)

	// e = e + 1
	e++
	fmt.Println(e)

	// e = e -1
	e--
	fmt.Println(e)
}