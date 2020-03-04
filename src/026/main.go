package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(cycles(1000))
}

func cycles(num int) int {
	arr := make([]string, num+1)
	arr[0] = ""
	arr[1] = ""
	for i := 2; i <= num; i++ {
		str := ""
		numb := 10
		mult := 10
		for j := 0; j < 10000; j++ {
			div := numb / i
			remainder := numb % i
			str += strconv.Itoa(div)
			if remainder == 0 {
				break
			} else {
				numb = remainder * mult
			}
		}
		arr[i] = str
	}

	maxRepeat := 1
	result := 2
	for i := 2; i < len(arr); i++ {
		if len(arr[i]) < 10000 {
			continue
		}
		for j := 0; j < len(arr[i]); j++ {
			testStr := arr[i]
			foundRepeat := false
			for k := 1; k < len(testStr)/2; k++ {
				testValue := testStr[j : j+k]
				values := len(testStr[j:]) / k
				count := strings.Count(testStr, testValue)
				if values == count && count > 2 {
					nextValue := testStr[j+k : j+k*2]
					if nextValue == testValue {
						if len(testValue) > maxRepeat {
							maxRepeat = len(testValue)
							result = i
						}
						foundRepeat = true
						break
					}
				}
			}
			if foundRepeat {
				break
			}
		}
	}

	return result
}
