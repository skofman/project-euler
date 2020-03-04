package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println(powers(100))
}

func powers(num int) int {
	arr := []*big.Int{}

	for a := 2; a <= num; a++ {
		for b := 2; b <= num; b++ {
			base, _ := new(big.Int).SetString(strconv.Itoa(a), 10)
			power, _ := new(big.Int).SetString(strconv.Itoa(b), 10)
			result := new(big.Int)
			result.Exp(base, power, result)
			if !doesInclude(arr, result) {
				arr = append(arr, result)
			}
		}
	}

	return len(arr)
}

func doesInclude(arr []*big.Int, result *big.Int) bool {
	for _, item := range arr {
		if item.Cmp(result) == 0 {
			return true
		}
	}
	return false
}
