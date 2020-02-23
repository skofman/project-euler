package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(getChain())
}

func getChain() int {
	result := 0

	for i := 2; i < 10000000; i++ {
		test := i
		for test != 1 && test != 89 {
			arr := strings.Split(strconv.Itoa(test), "")
			test = 0
			for _, s := range arr {
				num, _ := strconv.Atoi(s)
				test += num * num
			}
		}
		if test == 89 {
			result++
		}
	}

	return result
}
