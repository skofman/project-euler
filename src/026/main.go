package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(cycles(10))
}

func cycles(num int) int {
	arr := [][]int{}
	div := "10"
	score := 0
	for i := 1; i < num; i++ {
		emptyArr := []int{0, 0, 0, 0}
		arr = append(arr, emptyArr)
		if len(div) < len(strconv.Itoa(i)) {
			div += "0"
		}
		for len(arr[i-1] < 10000) {
			divNum, _ = strconv.Atoi(div)
			score = (int)(math.Floor((float64)(divNum) / (float64)(i)))
			arr[i-1] = append(arr[i-1], score)
			div = strconv.Itoa(divNum-score*i) + "0"
		}
		div = "10"
	}
	str := ""
	newArr := 
}
