package main

import (
	"fmt"
	"strconv"

	"../helpers"
)

func main() {
	fmt.Println(sumFactorial(100))
}

func sumFactorial(num int) int {
	result := "1"
	for i := num; i > 1; i-- {
		result = helpers.MultiplyBig(result, strconv.Itoa(i))
	}
	sum := 0
	for i := 0; i < len(result); i++ {
		intDigit, _ := strconv.Atoi(result[i : i+1])
		sum += intDigit
	}
	return sum
}
