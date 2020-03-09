package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println(getSum(99))
}

func getSum(num int) int {
	sum := 0
	arr := []int{2, 1}
	for i := 2; i <= num; i++ {
		if (i-2)%3 == 0 {
			arr = append(arr, 2+2*(i-2)/3)
		} else {
			arr = append(arr, 1)
		}
	}
	result := big.NewRat(1, int64(arr[len(arr)-1]))
	for i := len(arr) - 2; i > 0; i-- {
		rat := big.NewRat(int64(arr[i]), 1)
		result = rat.Add(rat, result)
		result = result.Inv(result)
	}
	result = result.Add(result, big.NewRat(2, 1))

	str := result.Num().Text(10)
	for _, val := range str {
		n, _ := strconv.Atoi(string(val))
		sum += n
	}
	return sum
}
