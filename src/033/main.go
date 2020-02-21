package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(curiousFractions())
}

func curiousFractions() float64 {
	var total float64 = 1
	var numerator float64
	var denom float64
	for i := 10; i <= 99; i++ {
		for j := i + 1; j <= 99; j++ {
			numeratorStr := strconv.Itoa(i)[0:1]
			denomStr := strconv.Itoa(j)[1:2]
			newNumerator, _ := strconv.Atoi(numeratorStr)
			newDenom, _ := strconv.Atoi(denomStr)
			numerator = (float64)(newNumerator)
			denom = (float64)(newDenom)
			if numerator/denom == (float64)(i)/(float64)(j) && strconv.Itoa(i)[1:2] == strconv.Itoa(j)[0:1] {
				total *= numerator
				total /= denom
			}
		}
	}
	return 1 / total
}
