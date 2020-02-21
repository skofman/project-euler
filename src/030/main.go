package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(powerFive())
}

func powerFive() int {
	sum := 0
	for i := 10; i < 1000000; i++ {
		arr := strings.Split(strconv.Itoa(i), "")
		power := 0
		for j := 0; j < len(arr); j++ {
			arrNum, _ := strconv.Atoi(arr[j])
			power += (int)(math.Pow((float64)(arrNum), (float64)(5)))
		}
		if power == i {
			sum += i
		}
	}
	return sum
}
