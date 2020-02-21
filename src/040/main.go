package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(champernowne())
}

func champernowne() int {
	d := 1
	i := 1
	counter := 0
	result := 1
	for d < 10000000 {
		str := strconv.Itoa(i)
		counter += len(str)
		if counter >= d {
			slice := d - counter - 1 + len(str)
			subStr := str[slice : slice+1]
			num, _ := strconv.Atoi(subStr)
			result *= num
			d *= 10
		}
		i++
	}
	return result
}
