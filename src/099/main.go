package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(findLargest())
}

func findLargest() int {
	content, _ := ioutil.ReadFile("./p099_base_exp.txt")
	numbers := string(content)
	rows := strings.Split(numbers, "\n")
	test := big.NewInt(1)
	line := 0

	for index, row := range rows {
		arr := strings.Split(row, ",")
		base, _ := strconv.Atoi(arr[0])
		baseInt := big.NewInt(int64(base))
		power, _ := strconv.Atoi(arr[1])
		powerInt := big.NewInt(int64(power))
		testInt := new(big.Int)
		testInt.Exp(baseInt, powerInt, nil)
		compare := testInt.Cmp(test)
		if compare == 1 {
			line = index + 1
			test = testInt
		}
	}

	return line
}
