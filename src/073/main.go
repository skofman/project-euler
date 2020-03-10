package main

import (
	//"../helpers"
	"fmt"
	"math"
)

func main() {
	fmt.Println(getFractionCount(12000))
}

func getFractionCount(num int) int {
	count := 0

	for i := num; i >= 1; i-- {
		start := int(math.Floor(float64(i) / 3))
		end := int(math.Ceil(float64(i) / 2))
		values := end - start - 1
		for j := start + 1; j < end; j++ {
			if j%2 == 0 && i%2 == 0 {
				values--
				continue
			}
			for k := 3; k < j; k += 2 {
				if j%k == 0 && i%k == 0 {
					values--
					break
				}
			}
		}
		count += values
	}

	return count
}
