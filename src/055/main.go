package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(findLychrel())
}

func findLychrel() int {
	count := 0
	palindrome := false
	for i := 5; i <= 10000; i++ {
		num := strconv.Itoa(i)
		for j := 0; j < 50; j++ {
			num = add(num)
			if isPalindrome(num) {
				palindrome = true
			}
		}
		if !palindrome {
			count++
		}
		palindrome = false
	}

	return count
}

func add(str string) string {
	arr1 := strings.Split(str, "")
	arr2 := []string{}
	for i := len(arr1) - 1; i >= 0; i-- {
		arr2 = append(arr2, arr1[i])
	}
	result := []string{}
	adder := 0
	for i := 0; i < len(arr1); i++ {
		num1, _ := strconv.Atoi(arr1[i])
		num2, _ := strconv.Atoi(arr2[i])
		temp := num1 + num2 + adder
		if temp > 9 {
			result = append(result, strings.Split(strconv.Itoa(temp), "")[1])
			adder = 1
		} else {
			result = append(result, strconv.Itoa(temp))
			adder = 0
		}
	}
	if adder == 1 {
		result = append(result, "1")
	}

	revResult := []string{}
	for i := len(result) - 1; i >= 0; i-- {
		revResult = append(revResult, result[i])
	}

	return strings.Join(revResult, "")
}

func isPalindrome(str string) bool {
	revStr := ""
	for _, s := range str {
		revStr = string(s) + revStr
	}

	return str == revStr
}
