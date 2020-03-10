package main

import (
	"../helpers"
	"fmt"
)

func main() {
	fmt.Println(getTotientMaximum(1000000))
}

func getTotientMaximum(num int) int {
	result := 1
	prime := 2

	for true {
		if result*prime > num {
			break
		}
		result *= prime
		prime = helpers.GetNextPrime(prime)
	}

	return result
}
