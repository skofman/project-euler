package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(findMinimumPath(80))
}

func findMinimumPath(size int) int {
	content, _ := ioutil.ReadFile("./p081_matrix.txt")
	numbers := string(content)
	rows := strings.Split(numbers, "\n")
	rows = rows[0 : len(rows)-1]
	arr := [][]string{}
	for _, row := range rows {
		arr = append(arr, strings.Split(row, ","))
	}

	result := [80][80]int{}
	firstNum, _ := strconv.Atoi(arr[0][0])
	result[0][0] = firstNum

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			num := result[i][j]
			if j != len(arr[i])-1 {
				nextNum, _ := strconv.Atoi(arr[i][j+1])
				if result[i][j+1] == 0 || result[i][j+1] > num+nextNum {
					result[i][j+1] = num + nextNum
				}
			}
			if i != size-1 {
				nextNum, _ := strconv.Atoi(arr[i+1][j])
				if result[i+1][j] == 0 || result[i+1][j] > num+nextNum {
					result[i+1][j] = num + nextNum
				}
			}
		}
	}

	return result[79][79]
}
