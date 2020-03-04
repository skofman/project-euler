package main

import (
	"fmt"
	"strconv"
	"strings"

	"../helpers"
)

func main() {
	fmt.Println(getChains())
}

func getChains() int {
	result := 0

	for i := 3; i < 1000000; i++ {
		arr := []int{i}
		for true {
			next := generateNext(arr[len(arr)-1])
			if arrContains(arr, next) {
				if len(arr) == 60 {
					result++
				}
				break
			} else if len(arr) == 60 {
				break
			} else {
				arr = append(arr, next)
			}
		}
	}

	return result
}

func arrContains(arr []int, val int) bool {
	for _, value := range arr {
		if value == val {
			return true
		}
	}

	return false
}

func generateNext(num int) int {
	str := strconv.Itoa(num)
	arr := strings.Split(str, "")
	sum := 0

	for _, val := range arr {
		value, _ := strconv.Atoi(val)
		sum += helpers.Factorial(value)
	}

	return sum
}
