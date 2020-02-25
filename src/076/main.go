package main

import (
	"fmt"
)

func main() {
	fmt.Println(findPermutations(100))
}

func findPermutations(num int) int {
	arr := []int{1, num - 1}
	result := 0

	for len(arr) > 0 {
		result++
		for len(arr) > 0 && arr[0] == 1 {
			arr = arr[1:]
		}
		if len(arr) == 0 {
			break
		}
		filler := arr[0] - 1
		arr = append([]int{filler}, arr[1:]...)
		for sumArr(arr) < num {
			if filler+sumArr(arr) <= num {
				arr = append([]int{filler}, arr...)
			} else {
				filler--
			}
		}
		fmt.Println(arr)
	}

	return result
}

func sumArr(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}

	return arr[0] + sumArr(arr[1:])
}
