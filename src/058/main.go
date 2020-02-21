package main

import (
	"fmt"

	"../helpers"
)

func main() {
	fmt.Println(spiralPrimes())
}

func spiralPrimes() int {
	walls := 1
	var ratio float64 = 1.0
	diagonalPrimes := []int{}
	nonPrimes := []int{}
	max := 1
	for ratio > 0.1 {
		walls += 2
		for i := max + walls - 1; i <= max+(walls-1)*4; i += (walls - 1) {
			if helpers.IsPrime(i) {
				diagonalPrimes = append(diagonalPrimes, i)
			} else {
				nonPrimes = append(nonPrimes, i)
			}
		}
		max = walls * walls
		ratio = float64(len(diagonalPrimes)) / (float64(len(nonPrimes)) + float64(len(diagonalPrimes)))
	}

	return walls
}
