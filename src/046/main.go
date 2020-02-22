package main

import (
	"fmt"

	"../helpers"
)

func main() {
	fmt.Println(findGoldbach())
}

func findGoldbach() int {
	primes := []int{2}
	i := 33
	latest := 33

	for true {
		for helpers.IsPrime(i) {
			i += 2
			latest = i
		}
		for primes[len(primes)-1] < i {
			primes = append(primes, getNextPrime(primes[len(primes)-1]))
		}
		for _, prime := range primes {
			j := 1
			calc := 0
			breakOut := false
			for calc <= i {
				calc = prime + 2*j*j
				if calc == i {
					i += 2
					j = 1
					calc = 0
					breakOut = true
					break
				}
				j++
			}
			if breakOut {
				break
			}
		}

		if latest == i {
			return i
		} else {
			latest = i
		}
	}

	return 0
}

func getNextPrime(num int) int {
	if helpers.IsPrime(num + 1) {
		return num + 1
	}

	return getNextPrime(num + 1)
}
