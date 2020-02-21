package main

import (
	"fmt"

	"../helpers"
)

func main() {
	fmt.Println(primeSum(2000000))
}

func primeSum(num int) int {
	sum := 5
	for i := 5; i < num; i++ {
		if helpers.IsPrime(i) {
			sum += i
		}
	}

	return sum
}
