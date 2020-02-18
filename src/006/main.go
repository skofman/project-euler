package main

import "fmt"

func main() {
	fmt.Println(findDiff(100))
}

func findDiff(num int) int {
	sum := 0
	squares := 0
	for i := 1; i <= num; i++ {
		sum += i
		squares += i * i
	}

	return sum*sum - squares
}
