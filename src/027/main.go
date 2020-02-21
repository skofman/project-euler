package main

import (
	"fmt"

	"../helpers"
)

func main() {
	fmt.Println(findPrimes())
}

func findPrimes() int {
	n := 0
	max := 0
	var result int
	for a := -999; a < 1000; a++ {
		for b := -999; b < 1000; b++ {
			for helpers.IsPrime(n*n + a*n + b) {
				n++
			}
			if n > max {
				max = n
				result = a * b
			}
			n = 0
		}
	}
	return result
}
