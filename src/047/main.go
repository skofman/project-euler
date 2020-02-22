package main

import (
	"fmt"

	"../helpers"
)

func main() {
	fmt.Println(distinctFactors())
}

func distinctFactors() int {
	primes := []int{2}
	i := 647

	for true {
		for primes[len(primes)-1] < i/2 {
			primes = append(primes, getNextPrime(primes[len(primes)-1]))
		}
		if testNumber(primes, i, 4) && testNumber(primes, i+1, 4) && testNumber(primes, i+2, 4) && testNumber(primes, i+3, 4) {
			return i
		}
		i++
	}

	return 0
}

func testNumber(primes []int, num int, count int) bool {
	factors := 0
	for _, prime := range primes {
		if num%prime == 0 {
			factors++
		}
		if count == factors {
			return true
		}
	}

	return false
}

func getNextPrime(num int) int {
	if helpers.IsPrime(num + 1) {
		return num + 1
	}

	return getNextPrime(num + 1)
}
