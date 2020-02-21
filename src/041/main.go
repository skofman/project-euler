package main

import (
	"fmt"
	"strconv"
	"strings"

	"../helpers"
)

func main() {
	fmt.Println(primePandigital())
}

func primePandigital() int {
	for i := 7654321; i > 0; i-- {
		if isPandigital(i) && helpers.IsPrime(i) {
			return i
		}
	}

	return 0
}

func isPandigital(num int) bool {
	arr := strings.Split(strconv.Itoa(num), "")
	max := 0
	for i := 0; i < len(arr); i++ {
		numArr, _ := strconv.Atoi(arr[i])
		if numArr > max {
			max = numArr
		}
	}

	if len(strconv.Itoa(num)) == len(helpers.UniqueStringSlice(arr)) && !helpers.SliceContainsString(arr, "0") && max == len(strconv.Itoa(num)) {
		return true
	}
	return false
}
