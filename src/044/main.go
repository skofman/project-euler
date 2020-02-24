package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(pentagonNumbers())
}

func pentagonNumbers() int {
	arr := []int{1, 5, 12}

	index := 2
	for true {
		for i := index; i >= 1; i-- {
			for j := i - 1; j >= 0; j-- {
				diff := arr[i] - arr[j]
				sum := arr[i] + arr[j]
				ind := sort.SearchInts(arr, diff)
				if sum >= arr[len(arr)-1] && arr[ind] == diff {
					for sum >= arr[len(arr)-1] {
						n := len(arr) + 1
						arr = append(arr, n*(3*n-1)/2)
					}
				}
				inx := sort.SearchInts(arr, sum)
				if arr[ind] == diff && arr[inx] == sum {
					return diff
				}
			}
		}
		index++
		if index == len(arr) {
			n := len(arr) + 1
			arr = append(arr, n*(3*n-1)/2)
		}
	}
	return 123456789
}
