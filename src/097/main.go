package main

import (
	"fmt"

	"../helpers"
)

func main() {
	fmt.Println(findTenDigits())
}

func findTenDigits() string {
	result := "2"
	for i := 1; i < 7830457; i++ {
		result = helpers.MultiplyBig(result, "2")
		if len(result) > 10 {
			result = result[len(result)-10:]
		}
	}
	result = helpers.MultiplyBig(result, "28433")
	if len(result) > 10 {
		result = result[len(result)-10:]
	}
	result = helpers.AddBigNum(result, "1")

	return result
}
