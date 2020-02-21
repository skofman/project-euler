package main

import (
	"fmt"
	"strconv"

	"../helpers"
)

func main() {
	fmt.Println(findCount())
}

func findCount() int {
	count := 0

	for i := 1; i <= 21; i++ {
		for j := 1; j <= 9; j++ {
			result := helpers.BigFactor(strconv.Itoa(j), i)
			if len(result) == i {
				count++
			}
		}
	}

	return count
}
