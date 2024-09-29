package main

import (
	"fmt"
	"math"
	"strconv"
)

func testNumber(str string, n int) bool {
	n1, _ := strconv.Atoi(str)
	if n1 == n {
		return true
	}

	for i := 1; i < len(str); i++ {
		n2, _ := strconv.Atoi(str[0:i])
		if testNumber(str[i:], n-n2) {
			return true
		}
	}

	return false
}

func isSNumber(n int) bool {
	result := math.Pow(float64(n), 2)
	str := strconv.Itoa(int(result))

	return testNumber(str, n)
}

func main() {
	maxNum := math.Pow(10, 12)
	sum := 0.0

	for i := 4; math.Pow(float64(i), 2) <= maxNum; i++ {
		if isSNumber(i) {
			num := math.Pow(float64(i), 2)
			fmt.Println(num, i)
			sum += num
		}
	}

	fmt.Printf("%d\n", int(sum))
}
