package main

import (
	"fmt"
	"strconv"

	"../helpers"
)

func main() {
	fmt.Println(powerfulDigits())
}

func powerfulDigits() int {
	result := 0

	for i := 1; i < 100; i++ {
		for j := 1; j < 100; j++ {
			temp := strconv.Itoa(i)
			for k := 1; k < j; k++ {
				temp = helpers.MultiplyBig(temp, strconv.Itoa(i))
			}
			fmt.Println(i, j)
			fmt.Println(temp)
			sum := sumDigits(temp)
			if sum > result {
				result = sum
			}
		}
	}
	return result
}

func sumDigits(str string) int {
	result := 0
	for _, s := range str {
		num, _ := strconv.Atoi(string(s))
		result += num
	}
	return result
}
