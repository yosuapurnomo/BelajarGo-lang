package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := sumAll(10, 5)
	fmt.Println(total)

	slice := []int{10, 15, total}

	sliceTotal := sumAll(slice...)
	fmt.Println(sliceTotal)
}
