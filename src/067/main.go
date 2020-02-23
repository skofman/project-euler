package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(maxPath())
}

func maxPath() float64 {
	content, _ := ioutil.ReadFile("./p067_triangle.txt")
	numbers := string(content)
	rows := strings.Split(numbers, "\n")
	strArr := [][]string{}
	for _, row := range rows {
		strArr = append(strArr, strings.Split(row, " "))
	}
	strArr = strArr[0 : len(strArr)-1]

	arr := [][]float64{}

	for i, row := range strArr {
		arr = append(arr, []float64{})
		for _, item := range row {
			num, _ := strconv.Atoi(item)
			arr[i] = append(arr[i], float64(num))
		}
	}

	for i := len(arr) - 2; i >= 0; i-- {
		for j := 0; j < len(arr[i]); j++ {
			arr[i][j] = arr[i][j] + math.Max(arr[i+1][j], arr[i+1][j+1])
		}
	}
	return arr[0][0]
}
