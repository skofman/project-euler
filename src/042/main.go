package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"../helpers"
)

func main() {
	fmt.Println(codedNumbers())
}

func codedNumbers() int {
	content, _ := ioutil.ReadFile("./p042_words.txt")
	words := string(content)
	wordsArr := strings.Split(words, ",")

	maxNumbers := []int{}
	for _, w := range wordsArr {
		sum := 0
		word := w[1 : len(w)-1]
		for j := 0; j < len(word); j++ {
			sum += int(word[j]) - 64
		}
		maxNumbers = append(maxNumbers, sum)
	}
	max := helpers.MaxInt(maxNumbers)

	triangles := []int{0}
	n := 1
	for triangles[n-1] < max {
		triangles = append(triangles, n*(n+1)/2)
		n++
	}
	count := 0
	for _, num := range maxNumbers {
		if helpers.SliceContainsInt(triangles, num) {
			count++
		}
	}
	return count
}
