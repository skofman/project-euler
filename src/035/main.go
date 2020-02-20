package main

import (
	"fmt"
	"strconv"
	"strings"

	"../helpers"
)

func main() {
	fmt.Println(circularPrimes(1000000))
}

func circularPrimes(num int) int {
	count := 0
	for i := 2; i < num; i++ {
		arr := strings.Split(strconv.Itoa(i), "")
		if primesConfirm(arr) {
			count++
		}
	}
	return count
}

func primesConfirm(arr []string) bool {
	for i := 0; i < len(arr); i++ {
		testNum, _ := strconv.Atoi(strings.Join(arr, ""))
		if !helpers.IsPrime(testNum) {
			return false
		}
		value := arr[0]
		arr = arr[1:]
		arr = append(arr, value)
	}
	return true
}
