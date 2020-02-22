package main

import (
	"fmt"
	"strconv"
	"strings"

	"../helpers"
)

func main() {
	fmt.Println(getSequencePrimes())
}

func getSequencePrimes() string {
	primes := []int{1489}

	for true {
		nextPrime := helpers.GetNextPrime(primes[len(primes)-1])
		if nextPrime >= 10000 {
			break
		}
		primes = append(primes, nextPrime)
	}

	for i, prime := range primes {
		for j := i + 1; j < len(primes); j++ {
			diff := primes[j] - prime
			thirdNum := primes[j] + diff
			if helpers.SliceContainsInt(primes, thirdNum) && arePermutations(prime, primes[j], thirdNum) {
				return strconv.Itoa(prime) + strconv.Itoa(primes[j]) + strconv.Itoa(primes[j]+diff)
			}
		}
	}

	return ""
}

func arePermutations(a int, b int, c int) bool {
	strA := strconv.Itoa(a)
	strB := strconv.Itoa(b)
	strC := strconv.Itoa(c)

	for i := 0; i < len(strA); i++ {
		count := strings.Count(strA, string(strA[i]))
		if count != strings.Count(strB, string(strA[i])) || count != strings.Count(strC, string(strA[i])) {
			return false
		}
	}

	return true
}
