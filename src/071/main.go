package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(getNumerator())
}

func getNumerator() *big.Int {
	min := big.NewRat(1, 8)
	max := big.NewRat(3, 7)

	for d := 3; d <= 1000000; d++ {
		start := d*3/7 - 1
		for n := start; n < d; n++ {
			test := big.NewRat(int64(n), int64(d))
			minCompare := min.Cmp(test)
			maxCompare := max.Cmp(test)
			if maxCompare == -1 {
				break
			}
			if minCompare == -1 && maxCompare == 1 {
				min = test
			}
		}
	}

	return min.Num()
}
