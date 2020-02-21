package main

import (
	"fmt"
	"strconv"
	"strings"

	"../helpers"
)

func main() {
	fmt.Println(digitFacts())
}

func digitFacts() int {
	arr := []string{}
	result := 0
	for i := 10; i < 362880; i++ {
		sum := 0
		arr = strings.Split(strconv.Itoa(i), "")
		for j := 0; j < len(arr); j++ {
			arrNum, _ := strconv.Atoi(arr[j])
			sum += helpers.Factorial(arrNum)
		}
		if sum == i {
			result += i
		}
	}
	return result
}
