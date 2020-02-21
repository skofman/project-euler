package main

import (
	"fmt"
	"strconv"

	"../helpers"
)

func main() {
	fmt.Println(orderNumbers(1000000, 10))
}

func orderNumbers(order int, digits int) string {
	arr := []int{}
	for i := 0; i < digits; i++ {
		arr = append(arr, i)
	}

	length := len(arr)
	result := ""
	for len(result) < length {
		permutations := helpers.Factorial(len(arr))
		index := (order - 1) / (permutations / len(arr))
		for index >= len(arr) {
			index -= len(arr)
		}
		result += strconv.Itoa(arr[index])
		arr = helpers.RemoveSliceIntElement(arr, index)
	}

	return result
}
