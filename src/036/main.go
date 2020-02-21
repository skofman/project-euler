package main

import (
	"fmt"
	"math"
	"strconv"

	"../helpers"
)

func main() {
	fmt.Println(doublePalindrome(1000000))
}

func doublePalindrome(num int) int {
	sum := 0
	for i := 1; i < num; i += 2 {
		str := strconv.Itoa(i)
		binary := makeBinary(i)
		if str == helpers.ReverseString(str) && binary == helpers.ReverseString(binary) {
			sum += i
		}
	}

	return sum
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
