package main

import "fmt"

func main() {
	fmt.Println(findPrime(6))
}

func findPrime(index int) int {
	arr := []int{2, 3}
	i := 5
	for len(arr) < index {
		if isPrime(i) {
			arr = append(arr, i)
		}
		i++
	}
	return arr[len(arr)-1]
}

func isPrime(num int) bool {
	return true
}
