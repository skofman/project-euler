package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findNextTriangle(285))
}

func findNextTriangle(num int) int {
	n := num + 1
	for true {
		triangle := n * (n + 1) / 2
		if isPentagonal(triangle) && isHexagonal(triangle) {
			return triangle
		}
		n++
	}

	return 0
}

func isPentagonal(num int) bool {
	a := float64(3)
	b := float64(-1)
	c := float64(-2 * num)
	result1, result2 := solveQuadratic(a, b, c)

	if result1 > 0 && isWhole(result1) {
		return true
	}
	if result2 > 0 && isWhole(result2) {
		return true
	}

	return false
}

func isHexagonal(num int) bool {
	a := float64(2)
	b := float64(-1)
	c := float64(-num)
	result1, result2 := solveQuadratic(a, b, c)

	if result1 > 0 && isWhole(result1) {
		return true
	}
	if result2 > 0 && isWhole(result2) {
		return true
	}

	return false
}

func solveQuadratic(a float64, b float64, c float64) (float64, float64) {
	result1 := (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)
	result2 := (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)

	return result1, result2
}

func isWhole(num float64) bool {
	_, frac := math.Modf(num)

	return frac == 0
}
