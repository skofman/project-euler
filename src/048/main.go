package main

import (
	"fmt"
	"strconv"

	"../helpers"
)

func main() {
	fmt.Println(findLastSum(1000, 10))
}

func findLastSum(num int, size int) string {
	sum := "0"
	for i := 1; i <= num; i++ {
		result := "1"
		for j := 1; j <= i; j++ {
			result = helpers.MultiplyBig(result, strconv.Itoa(i))
			if len(result) > size {
				index := len(result) - size
				result = result[index:]
			}
		}
		sum = helpers.AddBigNum(sum, result)
		if len(sum) > size {
			sum = sum[len(sum)-size:]
		}
	}

	return sum
}
