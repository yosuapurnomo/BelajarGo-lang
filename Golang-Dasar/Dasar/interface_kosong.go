package main

import "fmt"

func SayEverything(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Wowww"
	}
}

func main() {
	var say2 interface{} = SayEverything(2)
	fmt.Println(say2)
	var say1 interface{} = SayEverything(1)
	fmt.Println(say1)
	var say3 interface{} = SayEverything(3)
	fmt.Println(say3)
}
