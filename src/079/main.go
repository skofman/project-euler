package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println(findPasscode())
}

func findPasscode() string {
	content, _ := ioutil.ReadFile("./p079_keylog.txt")
	temp := strings.Split(string(content), "\n")

	logs := [][]string{}
	result := ""
	for _, row := range temp {
		if len(row) == 3 {
			split := strings.Split(row, "")
			logs = append(logs, split)
			for _, s := range row {
				if !strings.Contains(result, string(s)) {
					result += string(s)
				}
			}
		}
	}

	for _, code := range logs {
		if isSwap(result, code[0], code[1]) {
			i1 := strings.Index(result, code[0])
			i2 := strings.Index(result, code[1])
			arr := strings.Split(result, "")
			temp := arr[i1]
			arr[i1] = arr[i2]
			arr[i2] = temp
			result = strings.Join(arr, "")
		}
		if isSwap(result, code[0], code[2]) {
			i1 := strings.Index(result, code[0])
			i2 := strings.Index(result, code[2])
			arr := strings.Split(result, "")
			temp := arr[i1]
			arr[i1] = arr[i2]
			arr[i2] = temp
			result = strings.Join(arr, "")
		}
		if isSwap(result, code[1], code[2]) {
			i1 := strings.Index(result, code[1])
			i2 := strings.Index(result, code[2])
			arr := strings.Split(result, "")
			temp := arr[i1]
			arr[i1] = arr[i2]
			arr[i2] = temp
			result = strings.Join(arr, "")
		}
	}

	return result
}

func isSwap(str string, a string, b string) bool {
	i1 := strings.Index(str, a)
	i2 := strings.Index(str, b)

	return i2 < i1
}
