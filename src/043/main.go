package main

import (
	"fmt"
	"strconv"
	"strings"

	"../helpers"
)

func main() {
	// fmt.Println(subStringDiv())
	fmt.Println(getAllPandigitals([]int{1, 2, 0}))
}

func getAllPandigitals(s []int) []int {
	result := []int{}

	index := 0
	for index < len(s)-1 {
		startNum := s[index]
		result = append(result, generateNumber(s))
		nextNumber := moveNumber(s[index])
		for startNum != nextNumber {
			swapIndex := findIndex(s, nextNumber)
			swapNumbers(s, index, swapIndex)
			result = append(result, generateNumber(s))
			swapNumbers(s, index, swapIndex)
			nextNumber = moveNumber(nextNumber)
			if index == 0 && nextNumber == 0 {
				nextNumber = 1
			}
		}
		index++
	}

	return result
}

func swapNumbers(s []int, index int, swapIndex int) {
	temp := s[index]
	s[index] = s[swapIndex]
	s[swapIndex] = temp
}

func findIndex(s []int, num int) int {
	for i, item := range s {
		if item == num {
			return i
		}
	}

	return -1
}

func moveNumber(num int) int {
	if num == 2 {
		return 0
	}

	return num + 1
}

func generateNumber(s []int) int {
	strArr := []string{}
	for _, item := range s {
		strArr = append(strArr, strconv.Itoa(item))
	}
	strNum := strings.Join(strArr, "")
	num, _ := strconv.Atoi(strNum)

	return num
}

func subStringDiv() int {
	start := 1023456789
	end := 9876543210
	sum := 0
	for i := start; i <= end; i++ {
		if isPandigital(i) && isDiv(i) {
			sum += i
		}
	}

	return sum
}

func isDiv(num int) bool {
	strNum := strconv.Itoa(num)
	for i := 2; i <= 8; i++ {
		subStr := strNum[i : i+3]
		subNum, _ := strconv.Atoi(subStr)
		if !checkDiv(subNum) {
			return false
		}
	}
	return true
}

func checkDiv(num int) bool {
	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			return true
		}
	}
	return false
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
