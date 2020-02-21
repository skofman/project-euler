package main

import (
	"fmt"
	"strconv"

	"../helpers"
)

func main() {
	fmt.Println(truncatedPrimes())
}

func truncatedPrimes() int {
	count := 11
	i := 20
	sum := 0
	for count > 0 {
		if checkTruncation(i) {
			sum += i
			count--
		}
		i++
	}
	return sum
}

func checkTruncation(num int) bool {
	str := strconv.Itoa(num)
	length := len(str)
	arr := []string{}
	for i := 1; i < length; i++ {
		arr = append(arr, str[0:i], str[length-i:length])
	}
	arr = append(arr, str)
	for j := 0; j < len(arr); j++ {
		testValue, _ := strconv.Atoi(arr[j])
		if !helpers.IsPrime(testValue) {
			return false
		}
	}
	return true
}
