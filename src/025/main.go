package main

import (
	"fmt"

	"../helpers"
)

func main() {
	fmt.Println(fibDigits(3))
}

func fibDigits(num int) int {
	fib1 := "1"
	fib2 := "1"
	result := "1"
	index := 2

	for len(result) < 1000 {
		result = helpers.AddBigNum(fib1, fib2)
		fib1 = fib2
		fib2 = result
		index++
	}

	return index
}
