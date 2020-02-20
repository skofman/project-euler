package main

import (
	"fmt"
	"strconv"

	"../helpers"
)

func main() {
	fmt.Println(powers(100))
}

func powers(num int) int {
	arr := []string{}
	for a := 2; a <= num; a++ {
		for b := 2; b <= num; b++ {
			value := helpers.BigFactor(strconv.Itoa(a), b)
			if !helpers.SliceContainsString(arr, value) {
				arr = append(arr, value)
			}
		}
	}
	return len(arr)
}
