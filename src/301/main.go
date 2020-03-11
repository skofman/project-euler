package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getWins())
}

func getWins() int {
	count := 0

	max := int(math.Pow(2, 30))
	for i := 1; i <= max; i++ {
		if i^2*i^3*i == 0 {
			count++
		}
	}

	return count
}
