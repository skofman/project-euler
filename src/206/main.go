package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println(findInteger())
}

func findInteger() string {
	index := 0
	arr := [9]int{}

	for true {
		num := new(big.Float)
		strNum, _ := num.SetString(generateNum(arr))
		sqrt := num.Sqrt(strNum)
		if sqrt.IsInt() {
			return sqrt.Text('f', 0)
		}
		for true {
			if arr[index] == 9 {
				arr[index] = 0
				index++
			} else {
				arr[index]++
				index = 0
				break
			}
		}
	}

	return ""
}

func generateNum(arr [9]int) string {
	str := ""
	for i, num := range arr {
		str += strconv.Itoa(i + 1)
		str += strconv.Itoa(num)
	}
	str += "0"

	return str
}
