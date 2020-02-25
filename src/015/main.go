package main

import (
	"fmt"

	"../helpers"
)

func main() {
	fmt.Println(path(20))
}

func path(num int) int {
	rows := num + 1
	cols := num + 1
	row := []int{1}
	arr := [][]int{row}
	holder := make([]int, rows*cols)
	holder[0] = 1
	for len(arr) < 2*num+1 {
		newRow := []int{}
		for _, r := range arr[len(arr)-1] {
			if r%cols != 0 {
				if !helpers.SliceContainsInt(newRow, r+1) {
					newRow = append(newRow, r+1)
				}
				holder[r] += holder[r-1]

			}
			if r < num*rows+1 {
				if !helpers.SliceContainsInt(newRow, r+rows) {
					newRow = append(newRow, r+rows)
				}
				holder[r+rows-1] += holder[r-1]
			}
		}
		arr = append(arr, newRow)
	}

	return holder[len(holder)-1]
}
