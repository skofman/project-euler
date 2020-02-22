package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(pandigital())
}

func pandigital() int {
	start := 9999
	max := 0
	for i := start; i >= 1; i-- {
		j := 1
		test := ""
		for true {
			value := i * j
			test += strconv.Itoa(value)
			if len(test) > 9 || !isPandigital(test) {
				break
			} else if len(test) == 9 && isPandigital(test) {
				val, _ := strconv.Atoi(test)
				if val > max {
					max = val
					break
				}
			}
			j++
		}
	}

	return max
}

func isPandigital(str string) bool {
	if strings.Contains(str, "0") {
		return false
	}
	for _, s := range str {
		if strings.Count(str, string(s)) > 1 {
			return false
		}
	}

	return true
}
