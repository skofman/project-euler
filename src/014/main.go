package main

import "fmt"

func main() {
	fmt.Println(collatz(1000000))
}

func collatz(num int) int {
	max := 0
	result := 0
	chain := 0
	test := 0
	for i := num - 1; i > num/2; i-- {
		test = i
		for test != 1 {
			chain++
			if test%2 == 0 {
				test /= 2
			} else {
				test = 3*test + 1
			}
		}
		if chain > max {
			max = chain
			result = i
		}
		chain = 0
	}
	return result
}
