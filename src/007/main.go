package main

import (
	"fmt"

	"../helpers"
)

func main() {
	fmt.Println(findPrime(10001))
}

func findPrime(index int) int {
	arr := []int{2, 3}
	i := 5
	for len(arr) < index {
		if helpers.IsPrime(i) {
			arr = append(arr, i)
		}
		i++
	}
	return arr[len(arr)-1]
}
