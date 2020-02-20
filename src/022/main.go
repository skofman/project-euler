package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	fmt.Println(namesScore())
}

func namesScore() int {
	content, _ := ioutil.ReadFile("./p022_names.txt")
	names := string(content)
	namesArr := strings.Split(names, ",")
	sum := 0

	sort.Strings(namesArr)
	for i := 0; i < len(namesArr); i++ {
		name := namesArr[i][1 : len(namesArr[i])-1]
		for j := 0; j < len(name); j++ {
			sum += (int(name[j]) - 64) * (i + 1)
		}
	}

	return sum
}
