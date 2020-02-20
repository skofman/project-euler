package main

import (
	"fmt"

	"../helpers"
)

func main() {
	fmt.Println(abundant())
}

func findDivisorsSum(num int) int {
	sum := 0
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	return sum
}

func abundant() int {
	arr := []int{}

	for i := 2; i <= 28123; i++ {
		if findDivisorsSum(i) > i {
			arr = append(arr, i)
		}
	}
	sum := 0
	for i := 1; i <= 28123; i++ {
		if checkedOut(i, arr) {
			sum += i
		}
	}
	return sum
}

func checkedOut(num int, arr []int) bool {
	test := 0
	for i := 0; i < len(arr); i++ {
		if num < arr[i] {
			return true
		}
		test = num - arr[i]
		if helpers.SliceContainsInt(arr, test) {
			return false
		}
	}
	return true
}
