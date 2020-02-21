package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(product())
}

func numberIsPalindrome(num int) bool {
	str := strconv.Itoa(num)
	arr := reverseSlice(strings.Split(str, ""))
	revStr := strings.Join(arr, "")

	return str == revStr
}

func product() int {
	palindrome := 0
	for i := 100; i <= 999; i++ {
		for j := i; j <= 999; j++ {
			result := i * j
			if numberIsPalindrome(result) && result > palindrome {
				palindrome = result
			}
		}
	}
	return palindrome
}

func reverseSlice(slice []string) []string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
