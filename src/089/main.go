package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println(getRomanNumerals())
}

func getRomanNumerals() int {
	content, _ := ioutil.ReadFile("./p089_roman.txt")
	romans := strings.Split(string(content), "\n")

	sum := 0

	for _, val := range romans {
		sum += saveRomanCharacters(val)
	}

	return sum
}

func saveRomanCharacters(str string) int {
	sum := 0

	num := parseRomanToNumber(str)

	minimalRoman := parseNumberToRoman(num)

	sum += len(str) - len(minimalRoman)

	return sum
}

func parseNumberToRoman(num int) string {
	str := ""

	for num > 0 {
		if num >= 1000 {
			str += "M"
			num -= 1000
		} else if num >= 900 {
			str += "CM"
			num -= 900
		} else if num >= 500 {
			str += "D"
			num -= 500
		} else if num >= 400 {
			str += "CD"
			num -= 400
		} else if num >= 100 {
			str += "C"
			num -= 100
		} else if num >= 90 {
			str += "XC"
			num -= 90
		} else if num >= 50 {
			str += "L"
			num -= 50
		} else if num >= 40 {
			str += "XL"
			num -= 40
		} else if num >= 10 {
			str += "X"
			num -= 10
		} else if num >= 9 {
			str += "IX"
			num -= 9
		} else if num >= 5 {
			str += "V"
			num -= 5
		} else if num >= 4 {
			str += "IV"
			num -= 4
		} else {
			str += "I"
			num -= 1
		}
	}

	return str
}

func parseRomanToNumber(str string) int {
	sum := 0

	for i := 0; i < len(str); i++ {
		charStr := string(str[i])
		if i == len(str)-1 {
			sum += getRomanValue(charStr)
		} else if charStr == "M" || charStr == "D" || charStr == "L" || charStr == "V" {
			sum += getRomanValue(charStr)
		} else if charStr == "C" {
			nextStr := string(str[i+1])
			if nextStr == "D" || nextStr == "M" {
				sum += getRomanValue(nextStr) - getRomanValue(charStr)
				i++
			} else {
				sum += getRomanValue(charStr)
			}
		} else if charStr == "X" {
			nextStr := string(str[i+1])
			if nextStr == "L" || nextStr == "C" {
				sum += getRomanValue(nextStr) - getRomanValue(charStr)
				i++
			} else {
				sum += getRomanValue(charStr)
			}
		} else if charStr == "I" {
			nextStr := string(str[i+1])
			if nextStr == "V" || nextStr == "X" {
				sum += getRomanValue(nextStr) - getRomanValue(charStr)
				i++
			} else {
				sum += getRomanValue(charStr)
			}
		}
	}

	return sum
}

func getRomanValue(str string) int {
	switch str {
	case "M":
		return 1000
	case "D":
		return 500
	case "C":
		return 100
	case "L":
		return 50
	case "X":
		return 10
	case "V":
		return 5
	case "I":
		return 1
	}

	return 0
}
