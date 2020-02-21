package main

import (
	"fmt"

	"../helpers"
)

func main() {
	num := 600851475143
	fmt.Println(getLargePrime(num))
}

func getLargePrime(num int) int {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			div := num / i
			if helpers.IsPrime(div) {
				return div
			}
		}
	}

	return 1
}
