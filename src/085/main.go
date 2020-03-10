package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getGridArea(2000000))
}

func getGridArea(num int) int {
	diff := float64(num)
	area := 0
	for i := 1; i <= num; i++ {
		for j := 1; j <= num; j++ {
			rectangles := getRectangles(i, j, diff, num)
			if rectangles == -1 {
				break
			}
			if math.Abs(float64(rectangles-num)) < diff {
				diff = math.Abs(float64(rectangles - num))
				area = i * j
				if diff == 0 {
					return area
				}
			}
		}
	}

	return area
}

func getRectangles(x int, y int, diff float64, num int) int {
	rectangles := 0
	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			horizontal := x - i + 1
			vertical := y - j + 1
			rectangles += horizontal * vertical
			if rectangles > num+int(diff) {
				return -1
			}
		}
	}
	return rectangles
}
