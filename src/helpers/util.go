package helpers

import (
	"math"
	"strconv"
	"strings"
)

// GetNextPrime function
func GetNextPrime(num int) int {
	if IsPrime(num + 1) {
		return num + 1
	}

	return GetNextPrime(num + 1)
}

// MaxInt function
func MaxInt(arr []int) int {
	max := 0
	for _, item := range arr {
		if item > max {
			max = item
		}
	}
	return max
}

// MinInt function
func MinInt(arr []int) int {
	min := arr[0]
	for _, item := range arr {
		if item < min {
			min = item
		}
	}

	return min
}

// ReverseString function
func ReverseString(str string) string {
	arr := strings.Split(str, "")
	arr = ReverseSlice(arr)
	return strings.Join(arr, "")
}

// UniqueIntSlice function
func UniqueIntSlice(s []int) []int {
	arr := []int{}
	for _, item := range s {
		if !SliceContainsInt(arr, item) {
			arr = append(arr, item)
		}
	}

	return arr
}

// UniqueStringSlice function
func UniqueStringSlice(s []string) []string {
	arr := []string{}
	for _, item := range s {
		if !SliceContainsString(arr, item) {
			arr = append(arr, item)
		}
	}

	return arr
}

// RemoveSliceIntElement function
func RemoveSliceIntElement(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0
	s = s[:len(s)-1]

	return s
}

// Factorial function
func Factorial(num int) int {
	if num <= 1 {
		return 1
	}

	return num * Factorial(num-1)
}

// SliceContainsInt function
func SliceContainsInt(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}

	return false
}

// SliceContainsString function
func SliceContainsString(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}

	return false
}

// BigFactor function
func BigFactor(a string, b int) string {
	result := a
	for i := 1; i < b; i++ {
		result = MultiplyBig(result, a)
	}

	return result
}

// MultiplyBig function
func MultiplyBig(a string, b string) string {
	aa := ReverseSlice(strings.Split(a, ""))
	bb := ReverseSlice(strings.Split(b, ""))
	result := [][]int{}
	store := 0
	for i := 0; i < len(aa); i++ {
		if store > 0 {
			result[i-1] = append(result[i-1], store)
			store = 0
		}
		emptySlice := []int{}
		for k := 0; k < i; k++ {
			emptySlice = append(emptySlice, 0)
		}
		result = append(result, emptySlice)
		for j := 0; j < len(bb); j++ {
			aaNum, _ := strconv.Atoi(aa[i])
			bbNum, _ := strconv.Atoi(bb[j])
			mult := aaNum*bbNum + store
			if mult > 9 {
				multStr := strconv.Itoa(mult)
				newMult, _ := strconv.Atoi((string)(multStr[1]))
				newStore, _ := strconv.Atoi((string)(multStr[0]))
				result[i] = append(result[i], newMult)
				store = newStore
			} else {
				store = 0
				result[i] = append(result[i], mult)
			}
		}
		if store > 0 {
			result[i] = append(result[i], store)
			store = 0
		}
	}

	maxLength := 0
	for i := 0; i < len(result); i++ {
		if len(result[i]) > maxLength {
			maxLength = len(result[i])
		}
	}
	for i := 0; i < len(result); i++ {
		for j := len(result[i]); j < maxLength; j++ {
			result[i] = append(result[i], 0)
		}
	}

	if len(result) == 1 {
		str := strings.Join(ConvertSlice(result[0]), "")
		return strings.Join(ReverseSlice(strings.Split(str, "")), "")
	} else {
		str := strings.Join(ConvertSlice(result[0]), "")
		sum := strings.Join(ReverseSlice(strings.Split(str, "")), "")
		for i := 1; i < len(result); i++ {
			newStr := strings.Join(ConvertSlice(result[i]), "")
			sum = AddBigNum(sum, strings.Join(ReverseSlice(strings.Split(newStr, "")), ""))
		}
		return sum
	}
}

// AddBigNum function
func AddBigNum(a string, b string) string {
	aa := ReverseSlice(strings.Split(a, ""))
	bb := ReverseSlice(strings.Split(b, ""))
	result := []string{}
	store := 0
	if len(aa) > len(bb) {
		for i := len(bb); i < len(aa); i++ {
			bb = append(bb, "0")
		}
	} else if len(bb) > len(aa) {
		for i := len(aa); i < len(bb); i++ {
			aa = append(aa, "0")
		}
	}

	for i := 0; i < len(aa); i++ {
		aaNum, _ := strconv.Atoi(aa[i])
		bbNum, _ := strconv.Atoi(bb[i])
		sum := aaNum + bbNum + store
		if sum > 9 {
			sumStr := strconv.Itoa(sum)
			result = append(result, sumStr[1:2])
			newStore, _ := strconv.Atoi(sumStr[0:1])
			store = newStore
		} else {
			result = append(result, strconv.Itoa(sum))
			store = 0
		}
	}

	if store > 0 {
		result = append(result, strconv.Itoa(store))
	}

	return strings.Join(ReverseSlice(result), "")
}

// ConvertSlice int to string
func ConvertSlice(s []int) []string {
	result := []string{}
	for i := 0; i < len(s); i++ {
		result = append(result, strconv.Itoa(s[i]))
	}

	return result
}

// IsPrime function
func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	testValue := (int)(math.Round(math.Sqrt((float64)(num))))
	for i := 2; i <= testValue; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

// ReverseSlice function
func ReverseSlice(slice []string) []string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
