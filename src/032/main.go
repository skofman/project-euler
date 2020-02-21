package main

import (
	"fmt"
	"strconv"
	"strings"

	"../helpers"
)

func main() {
	fmt.Println(pandigital())
}

func pandigital() int {
	sum := 0
	multiply := 0
	str := ""
	arr := []string{}
	result := []int{}
	for i := 1; i < 5000; i++ {
		for j := i; j < 5000; j++ {
			multiply = i * j
			str = strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(multiply)
			arr = helpers.UniqueStringSlice(strings.Split(str, ""))
			if len(arr) == 9 && len(str) == 9 && !helpers.SliceContainsString(arr, "0") {
				result = append(result, multiply)
			}
		}
	}
	result = helpers.UniqueIntSlice(result)
	for i := 0; i < len(result); i++ {
		sum += result[i]
	}
	return sum
}
