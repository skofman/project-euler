package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(findNumber())
}

func findNumber() int {
	found := true

	for i := 125873; i < 10000000; i++ {
		for j := 2; j <= 6; j++ {
			test := i * j
			a := strings.Split(strconv.Itoa(i), "")
			b := strings.Split(strconv.Itoa(test), "")
			sort.Strings(a)
			sort.Strings(b)
			for k := 0; k < len(a); k++ {
				if a[k] != b[k] {
					found = false
					break
				}
			}
		}

		if found {
			return i
		} else {
			found = true
		}

		newA := strings.Split(strconv.Itoa(i), "")
		if newA[0] != "1" {
			length := len(strconv.Itoa(i))
			i = int(math.Pow(10, float64(length)))
		}
	}

	return -1
}
