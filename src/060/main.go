package main

import (
	"fmt"
	"strconv"

	"../helpers"
)

func main() {
	fmt.Println(findSum())
}

func findSum() int {
	sumFound := false
	primes := []int{3}
	testingPrime := getNextPrime(3)

	for !sumFound {
		if len(primes) == 1 {
			if checkTwoNumbers(primes[0], testingPrime) {
				primes = append(primes, testingPrime)
			}
		} else if len(primes) == 2 {
			if checkTwoNumbers(primes[0], testingPrime) && checkTwoNumbers(primes[1], testingPrime) {
				primes = append(primes, testingPrime)
			}
		} else if len(primes) == 3 {
			if checkTwoNumbers(primes[0], testingPrime) &&
				checkTwoNumbers(primes[1], testingPrime) &&
				checkTwoNumbers(primes[2], testingPrime) {
				primes = append(primes, testingPrime)
			}
		} else if len(primes) == 4 {
			if checkTwoNumbers(primes[0], testingPrime) &&
				checkTwoNumbers(primes[1], testingPrime) &&
				checkTwoNumbers(primes[2], testingPrime) &&
				checkTwoNumbers(primes[3], testingPrime) {
				primes = append(primes, testingPrime)
			}
		}

		if len(primes) == 5 {
			sumFound = true
		} else if len(primes) == 0 {
			fmt.Println("Increase test")
			sumFound = true
		}
		if testingPrime > 10000 {
			if len(primes) == 1 {
				primes = []int{getNextPrime(primes[0])}
				testingPrime = primes[0]
			} else {
				testingPrime = getNextPrime(primes[len(primes)-1])
				primes = primes[0 : len(primes)-1]
			}
		} else {
			testingPrime = getNextPrime(testingPrime)
		}
	}

	sum := 0
	for _, prime := range primes {
		sum += prime
	}

	return sum
}

func getNextPrime(num int) int {
	i := num + 1
	if helpers.IsPrime(i) {
		return i
	} else {
		return getNextPrime(i)
	}
}

func checkTwoNumbers(a int, b int) bool {
	strA := strconv.Itoa(a) + strconv.Itoa(b)
	strB := strconv.Itoa(b) + strconv.Itoa(a)
	numA, _ := strconv.Atoi(strA)
	numB, _ := strconv.Atoi(strB)
	return helpers.IsPrime(numA) && helpers.IsPrime(numB)
}
