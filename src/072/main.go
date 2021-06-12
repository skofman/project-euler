package main

import (
	"fmt"
	"project-euler/src/helpers"
)

func main() {
	fmt.Println(getNumberOfElements(1000000))
}

func getNumberOfElements(max int) int {
	sum := 1
	fractions := map[int]int{}
	fractions[2] = 1
	sorted := []int{2}

	primes := []int{}

	for i := max; i >= 2; i-- {
		if helpers.IsPrime(i) {
			primes = append(primes, i)
		}
	}

	for d := 3; d <= max; d++ {
		if primes[0] == d {
			fractions[d] = d - 1
			sum += d - 1
			sorted = append(sorted, d)
			primes = primes[1:]
		} else {
			temp := d - 1
			for index, key := range sorted {
				if d%key == 0 {
					temp -= fractions[key]
				}
				if index > len(sorted)/2 {
					break
				}
			}
			sum += temp
			fractions[d] = temp
			sorted = append(sorted, d)
		}
	}

	return sum
}
