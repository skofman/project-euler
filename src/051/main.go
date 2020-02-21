package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"../helpers"
)

func main() {
	fmt.Println(primeDigitReplacements())
}

func primeDigitReplacements() int {
	solutionFound := false
	prime := 101
	primes := 0
	nonPrimes := 0
	primeList := []string{}

	for !solutionFound {
		length := len(strconv.Itoa(prime)) - 1
		end := int(math.Pow(2, float64(length)))
		num := strings.Split(strconv.Itoa(prime), "")

		for j := 1; j < end; j++ {
			test := makeBinary(j)
			switch length - len(test) {
			case 1:
				test = "0" + test
			case 2:
				test = "00" + test
			case 3:
				test = "000" + test
			case 4:
				test = "0000" + test
			}
			for i := 0; i < 10; i++ {
				testNumber := make([]string, len(num))
				copy(testNumber, num)
				for k := 0; k < len(test); k++ {
					if test[k:k+1] == "1" {
						testNumber[k] = strconv.Itoa(i)
					}
				}
				numTest, _ := strconv.Atoi(strings.Join(testNumber, ""))
				if helpers.IsPrime(numTest) && testNumber[0] != "0" {
					primes++
					primeList = append(primeList, strings.Join(testNumber, ""))
				} else {
					nonPrimes++
				}
				if nonPrimes > 2 {
					break
				}
			}
			if primes == 8 {
				numList := []int{}
				for _, item := range primeList {
					num, _ := strconv.Atoi(item)
					numList = append(numList, num)
				}
				sort.Ints(numList)
				return numList[0]
				solutionFound = true
			} else {
				primes = 0
				nonPrimes = 0
				primeList = []string{}
			}
		}
		prime = getNextPrime(prime)
	}

	return prime
}

func getNextPrime(num int) int {
	i := num + 1
	if helpers.IsPrime(i) {
		return i
	} else {
		return getNextPrime(i)
	}
}

func makeBinary(num int) string {
	str := ""
	for num > 0 {
		if num%2 == 0 {
			str += "0"
		} else {
			str += "1"
		}
		num = (int)(math.Floor((float64)(num) / 2.0))
	}
	return str
}
