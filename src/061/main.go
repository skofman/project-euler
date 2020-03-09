package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(getSum())
}

type series struct {
	used     [6]int
	sequence []int
}

func getSum() int {
	variations := [6][]int{generateArray("triangle"), generateArray("square"), generateArray("pentagonal"), generateArray("hexagonal"), generateArray("heptagonal"), generateArray("octagonal")}
	var result []series
	sum := 0

	for len(result) != 1 {
		if len(result) == 0 {
			for _, triangle := range variations[0] {
				usedArr := [6]int{}
				usedArr[0] = 1
				seq := []int{triangle}
				result = append(result, series{usedArr, seq})
			}
		}
		var newResult []series
		for _, val := range result {
			for i := 0; i < len(val.used); i++ {
				if val.used[i] == 0 {
					seq := make([]int, len(val.sequence))
					copy(seq, val.sequence)
					numbers := findNumbers(seq[len(seq)-1], variations[i])
					for _, num := range numbers {
						seq = append(seq, num)
						newUsed := val.used
						newUsed[i] = 1
						newResult = append(newResult, series{newUsed, seq})
					}
				}
			}
		}
		result = newResult
		if len(result[0].sequence) == 6 {
			break
		}
	}

	for _, val := range result {
		first := strconv.Itoa(val.sequence[0])[:2]
		last := strconv.Itoa(val.sequence[5])[2:]
		if first == last {
			for _, i := range val.sequence {
				sum += i
			}
		}
	}

	return sum
}

func findNumbers(num int, arr []int) []int {
	var result []int
	str := strconv.Itoa(num)[2:]

	for _, item := range arr {
		test := strconv.Itoa(item)[:2]
		if test == str {
			result = append(result, item)
		}
	}

	return result
}

func generateArray(types string) []int {
	n := 1
	num := 0
	var arr []int
	for num < 10000 {
		switch types {
		case "triangle":
			num = n * (n + 1) / 2
		case "square":
			num = n * n
		case "pentagonal":
			num = n * (3*n - 1) / 2
		case "hexagonal":
			num = n * (2*n - 1)
		case "heptagonal":
			num = n * (5*n - 3) / 2
		case "octagonal":
			num = n * (3*n - 2)
		}
		if num >= 1000 && num < 10000 {
			arr = append(arr, num)
		}
		n++
	}

	return arr
}
