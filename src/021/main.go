package main

import "fmt"

func main() {
	fmt.Println(amicable(10000))
}

func amicable(num int) int {
	sum := 0
	test := 0
	for i := 1; i < num; i++ {
		test = findDivisorsSum(i)
		if test != i && test != 1 && i == findDivisorsSum(test) {
			sum += i
		}
	}

	return sum
}

func findDivisorsSum(num int) int {
	sum := 0
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			sum += i
		}
	}

	return sum
}
