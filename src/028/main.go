package main

import "fmt"

func main() {
	fmt.Println(spiral(1001))
}

func spiral(num int) int {
	sum := 1
	prevRight := 1
	prevLeft := 1
	testValue := num / 2
	for i := 1; i <= testValue; i++ {
		sum += (2*i + 1) * (2*i + 1)
		sum += (4*i*i + 1)
		sum += prevRight + 8*i - 6
		prevRight += 8*i - 6
		sum += prevLeft + 8*i - 2
		prevLeft += 8*i - 2
	}

	return sum
}
