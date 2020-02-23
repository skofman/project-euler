package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(getSum())
}

func getSum() int {
	content, _ := ioutil.ReadFile("./p059_cipher.txt")
	codes := string(content)
	arr := strings.Split(codes, ",")
	numbers := []int{}

	for _, str := range arr {
		val, _ := strconv.Atoi(str)
		numbers = append(numbers, val)
	}

	// findKey(numbers)
	// Found key = exp
	key := "exp"
	sum := 0
	index := 0
	for _, num := range numbers {
		sum += num ^ int(key[index])
		if index == 2 {
			index = 0
		} else {
			index++
		}
	}

	return sum
}

func findKey(numbers []int) {
	for i := 97; i <= 122; i++ {
		for j := 97; j <= 122; j++ {
			for k := 97; k <= 122; k++ {
				key := string(i) + string(j) + string(k)
				strArr := []string{}
				index := 0
				for _, num := range numbers {
					result := num ^ int(key[index])
					strArr = append(strArr, string(result))
					if index == 2 {
						index = 0
					} else {
						index++
					}
				}
				str := strings.Join(strArr, "")
				if strings.Contains(str, "how") {
					fmt.Println(str)
					fmt.Println(key)
					reader := bufio.NewReader(os.Stdin)
					text, _ := reader.ReadString('\n')
					fmt.Println(text)
				}
			}
		}
	}
}
