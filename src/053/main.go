package main

import (
	"fmt"

	"../helpers"
)

func main() {
	fmt.Println(findValues())
}

func findValues() int {
	count := 0

	for i := 1; i <= 100; i++ {
		for j := 1; j <= i; j++ {
			value := float64(helpers.Factorial(i)) / (float64(helpers.Factorial(j)) * float64(helpers.Factorial(i-j)))
			if value > 1000000.0 {
				count++
			}
		}
	}

	return count
}
