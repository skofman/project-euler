package main

import "fmt"

func main() {
	fmt.Println(pyth())
}

func pyth() int {
	c := 0
	for a := 1; a < 1000; a++ {
		for b := a + 1; b < 1000; b++ {
			c = 1000 - b - a
			if c*c == a*a+b*b {
				return a * b * c
			}
		}
	}
	return 1
}
