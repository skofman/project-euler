package main

import "fmt"

func main() {
	fmt.Println(findTriangles())
}

func findTriangles() int {
	p := 0
	max := 0
	for i := 3; i <= 1000; i++ {
		counter := 0
		for a := 1; a <= 1000; a++ {
			for b := a + 1; b <= 1000; b++ {
				c := i - b - a
				if a*a+b*b == c*c && c > 0 {
					counter++
				}
			}
		}
		if counter > max {
			max = counter
			p = i
		}
	}
	return p
}
