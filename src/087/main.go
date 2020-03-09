package main

import (
	"../helpers"
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(getPrimeTriples(50000000))
}

func getPrimeTriples(max int) int {
	var result []int
	maxPrime, _ := math.Modf(math.Sqrt(float64(max)))
	primes := []int{2}

	for i := 3; i <= int(maxPrime)+1; i++ {
		if helpers.IsPrime(i) {
			primes = append(primes, i)
		}
	}

	for _, a := range primes {
		for _, b := range primes {
			for _, c := range primes {
				calc := math.Pow(float64(a), 2) + math.Pow(float64(b), 3) + math.Pow(float64(c), 4)
				testResult := int(calc)
				index := sort.SearchInts(result, testResult)
				if testResult < max {
					if index == len(result) {
						result = append(result, testResult)
					} else if testResult != result[index] {
						result = append(result[:index], append([]int{testResult}, result[index:]...)...)
					}
				} else {
					break
				}
			}
			if math.Pow(float64(a), 2)+math.Pow(float64(b), 3) > float64(max) {
				break
			}
		}
	}
	return len(result)
}
