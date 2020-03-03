package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(subStringDiv())
}

func getAllPandigitals(n int, A []int) []string {
	result := []string{}
	c := []int{}

	for i := 0; i < n; i++ {
		c = append(c, 0)
	}

	str := generateString(A)
	result = append(result, str)

	i := 0
	for i < n {
		if c[i] < i {
			if i%2 == 0 {
				A = swap(A, 0, i)
			} else {
				A = swap(A, c[i], i)
			}
			str := generateString(A)
			result = append(result, str)
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}

	return result
}

func generateString(arr []int) string {
	str := ""

	for _, num := range arr {
		str += strconv.Itoa(num)
	}

	return str
}

func subStringDiv() int {
	numbers := getAllPandigitals(10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	primes := []int{2, 3, 5, 7, 11, 13, 17}
	sum := 0

	for _, num := range numbers {
		for i := 1; i <= 7; i++ {
			testNum, _ := strconv.Atoi(num[i : i+3])
			div := true
			if testNum%primes[i-1] != 0 {
				div = false
			}
			if !div {
				break
			}
			if i == 7 {
				number, _ := strconv.Atoi(num)
				sum += number
			}
		}

	}

	return sum
}

func swap(arr []int, i1 int, i2 int) []int {
	temp := arr[i1]
	arr[i1] = arr[i2]
	arr[i2] = temp

	return arr
}
