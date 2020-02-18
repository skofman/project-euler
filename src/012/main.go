package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(triangle(500))
}

func triangle(num int) int {
	arr := []int{1}
	var divisors int

	for divisors < num {
		arr = append(arr, arr[len(arr)-1]+len(arr)+1)
		value := math.Sqrt((float64)(arr[len(arr)-1])) * math.Sqrt((float64)(arr[len(arr)-1]))
		if (float64)(arr[len(arr)-1]) == value {
			divisors = 1
			for i := 1; i < (int)(math.Sqrt((float64)(arr[len(arr)-1]))); i++ {
				if arr[len(arr)-1]%i == 0 {
					divisors += 2
				}
			}
		}
	}
	return arr[len(arr)-1]
}
