package main

import (
	"fmt"
)

func main() {
	fmt.Println(findValues())
}

func findValues() int {
	total := 0

	for i := 10; i <= 100; i++ {
		for j := 1; j <= i; j++ {
			if isMill(i, j) {
				total++
			}
		}
	}

	return total
}

func isMill(n int, r int) bool {
	diff := n - r
	if diff == 0 {
		return false
	}
	startTop := float64(n)
	endTop := float64(r)
	startBottom := float64(n - r)
	endBottom := float64(1)
	result := float64(1)
	action := "*"
	for true {
		if startTop == endTop && result <= 1000000 {
			return false
		}
		if startBottom == endBottom && result > 1000000 {
			return true
		}
		if action == "*" {
			result *= startTop
			startTop--
			if result > 1000000 {
				action = "/"
			}
		} else if action == "/" {
			result /= startBottom
			startBottom--
			if result <= 1000000 {
				action = "*"
			}
		}

	}

	return true
}
