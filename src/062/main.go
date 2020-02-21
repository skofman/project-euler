package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(cubicPermutation())
}

func cubicPermutation() int {
	max := 208064
	i := 1
	for i < max {
		numbers := []int{}
		cubic := cube(i)
		for j := i + 1; j < max; j++ {
			testNumber := cube(j)
			if len(strconv.Itoa(testNumber)) > len(strconv.Itoa(cubic)) {
				break
			}
			if isPermutation(cubic, testNumber) {
				numbers = append(numbers, testNumber)
				if len(numbers) == 4 {
					return cubic
				}
			}
		}
		i++
	}
	return 0
}

func cube(num int) int {
	return num * num * num
}

func isPermutation(a int, b int) bool {
	arrA := strings.Split(strconv.Itoa(a), "")
	arrB := strings.Split(strconv.Itoa(b), "")
	sort.Sort(sort.StringSlice(arrA))
	sort.Sort(sort.StringSlice(arrB))

	return strings.Join(arrA, "") == strings.Join(arrB, "")
}
