package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findBouncy())
}

func findBouncy() int {
	nonBouncy := 99.0
	bouncy := 0.0
	num := 99
	for bouncy/(bouncy+nonBouncy) != 0.99 {
		num++
		if isBouncy(num) {
			bouncy++
		} else {
			nonBouncy++
		}
	}

	return num
}

func isBouncy(num int) bool {
	str := strconv.Itoa(num)
	direction := 0
	for i := 1; i < len(str); i++ {
		next, _ := strconv.Atoi(string(str[i]))
		prev, _ := strconv.Atoi(string(str[i-1]))
		if direction == 1 && prev > next {
			return true
		} else if direction == -1 && prev < next {
			return true
		} else if direction == 0 {
			if prev < next {
				direction = 1
			} else if next < prev {
				direction = -1
			}
		}
	}

	return false
}
