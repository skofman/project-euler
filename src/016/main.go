package main

import (
	"fmt"
	"strconv"

	"../helpers"
)

func main() {
	fmt.Println(sumDigits(1000))
}

func sumDigits(num int) int {
	resultStr := helpers.BigFactor("2", num)
	sum := 0
	for i := 0; i < len(resultStr); i++ {
		j, _ := strconv.Atoi(resultStr[i : i+1])
		sum += j
	}
	return sum
}
