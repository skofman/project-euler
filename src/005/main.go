package main

import "fmt"

func main() {
	fmt.Println(dividers(20))
}

func dividers(num int) int {
	test := num
	for i := num; i > 0; i-- {
		if test%i != 0 {
			test += num
			i = num
		}
	}

	return test
}
