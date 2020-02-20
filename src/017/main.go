package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(letterCount(1000))
}

func letterCount(num int) int {
	sum := 0
	for i := 1; i < num; i++ {
		sum += convert(i)
	}
	sum += len("onethousand")
	return sum
}

func convert(num int) int {
	arr := strings.Split(strconv.Itoa(num), "")
	str := ""
	if len(arr) == 1 {
		switch arr[0] {
		case "1":
			str += "one"
		case "2":
			str += "two"
		case "3":
			str += "three"
		case "4":
			str += "four"
		case "5":
			str += "five"
		case "6":
			str += "six"
		case "7":
			str += "seven"
		case "8":
			str += "eight"
		case "9":
			str += "nine"
		}
	}
	if len(arr) == 2 {
		if arr[0] == "1" {
			switch arr[1] {
			case "0":
				str += "ten"

			case "1":
				str += "eleven"

			case "2":
				str += "twelve"

			case "3":
				str += "thirteen"

			case "4":
				str += "fourteen"

			case "5":
				str += "fifteen"

			case "6":
				str += "sixteen"

			case "7":
				str += "seventeen"

			case "8":
				str += "eighteen"

			case "9":
				str += "nineteen"

			}
		} else {
			switch arr[1] {
			case "1":
				str += "one"

			case "2":
				str += "two"

			case "3":
				str += "three"

			case "4":
				str += "four"

			case "5":
				str += "five"

			case "6":
				str += "six"

			case "7":
				str += "seven"

			case "8":
				str += "eight"

			case "9":
				str += "nine"

			}
			switch arr[0] {
			case "2":
				str += "twenty"

			case "3":
				str += "thirty"

			case "4":
				str += "forty"

			case "5":
				str += "fifty"

			case "6":
				str += "sixty"

			case "7":
				str += "seventy"

			case "8":
				str += "eighty"

			case "9":
				str += "ninety"

			}
		}
	}
	if len(arr) == 3 {
		switch arr[0] {
		case "1":
			str += "onehundred"

		case "2":
			str += "twohundred"

		case "3":
			str += "threehundred"

		case "4":
			str += "fourhundred"

		case "5":
			str += "fivehundred"

		case "6":
			str += "sixhundred"

		case "7":
			str += "sevenhundred"

		case "8":
			str += "eighthundred"

		case "9":
			str += "ninehundred"

		}
		switch arr[1] {
		case "2":
			str += "andtwenty"

		case "3":
			str += "andthirty"

		case "4":
			str += "andforty"

		case "5":
			str += "andfifty"

		case "6":
			str += "andsixty"

		case "7":
			str += "andseventy"

		case "8":
			str += "andeighty"

		case "9":
			str += "andninety"

		}
		if arr[1] == "1" {
			switch arr[2] {
			case "0":
				str += "andten"

			case "1":
				str += "andeleven"

			case "2":
				str += "andtwelve"

			case "3":
				str += "andthirteen"

			case "4":
				str += "andfourteen"

			case "5":
				str += "andfifteen"

			case "6":
				str += "andsixteen"

			case "7":
				str += "andseventeen"

			case "8":
				str += "andeighteen"

			case "9":
				str += "andnineteen"

			}
		} else if arr[1] == "0" {
			switch arr[2] {
			case "1":
				str += "andone"

			case "2":
				str += "andtwo"

			case "3":
				str += "andthree"

			case "4":
				str += "andfour"

			case "5":
				str += "andfive"

			case "6":
				str += "andsix"

			case "7":
				str += "andseven"

			case "8":
				str += "andeight"

			case "9":
				str += "andnine"

			}
		} else {
			switch arr[2] {
			case "1":
				str += "one"

			case "2":
				str += "two"

			case "3":
				str += "three"

			case "4":
				str += "four"

			case "5":
				str += "five"

			case "6":
				str += "six"

			case "7":
				str += "seven"

			case "8":
				str += "eight"

			case "9":
				str += "nine"

			}
		}
	}
	return len(str)
}
