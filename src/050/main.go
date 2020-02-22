package main

import (
	"fmt"

	"../helpers"
)

func main() {
	fmt.Println(consecutivePrimeSum())
}

func consecutivePrimeSum() int {
	primes := []int{2}
	sequence := 0
	prime := 0
	for true {
		nextPrime := helpers.GetNextPrime(primes[len(primes)-1])
		if nextPrime > 1000000 {
			break
		}
		primes = append(primes, nextPrime)
	}

	for i := 0; i < len(primes); i++ {
		sum := 0
		count := 0
		for j := i; j < len(primes); j++ {
			sum += primes[j]
			count++
			if sum >= 1000000 {
				break
			}
			if helpers.SliceContainsInt(primes, sum) && count > sequence {
				sequence = count
				prime = sum
			}
		}
	}

	return prime
}
