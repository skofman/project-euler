package main

import "fmt"

func main() {
	i := 1
	j := 2
	sum := 2

	for j <= 4000000 {
		k := i + j
		i = j
		j = k
		if k%2 == 0 {
			sum += k
		}
	}

	fmt.Println(sum)
}
