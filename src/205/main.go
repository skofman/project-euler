package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getProbability())
}

func getProbability() float64 {
	peter := make(map[int]int)
	colin := make(map[int]int)

	totalPeter := math.Pow(4.0, 9.0)
	totalColin := math.Pow(6.0, 6.0)

	fourSide := []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	sixSide := []int{1, 1, 1, 1, 1, 1}

	for peter[36] == 0 {
		sum := sumSlice(fourSide)
		fourSide = getNextSlice(fourSide, 4)
		peter[sum]++
	}
	for colin[36] == 0 {
		sum := sumSlice(sixSide)
		sixSide = getNextSlice(sixSide, 6)
		colin[sum]++
	}

	total := 0.0
	for key, val := range peter {
		peteValue := key
		peterOdds := float64(val) / totalPeter
		colinSum := 0
		for key, val := range colin {
			if key < peteValue {
				colinSum += val
			}
		}
		colinOdds := float64(colinSum) / totalColin
		total += peterOdds * colinOdds
	}

	return total
}

func getNextSlice(arr []int, max int) []int {
	var result []int
	size := len(arr)
	index := 0

	for i := 0; i < size; i++ {
		if arr[i] == max {
			index = i + 1
		} else {
			break
		}
	}

	for i := 0; i < index; i++ {
		result = append(result, 1)
	}
	if len(result) < size {
		result = append(result, arr[index]+1)
	}
	if len(result) < size {
		result = append(result, arr[index+1:]...)
	}

	return result
}

func sumSlice(arr []int) int {
	total := 0
	for _, val := range arr {
		total += val
	}

	return total
}
