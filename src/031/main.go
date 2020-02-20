package main

import "fmt"

func main() {
	fmt.Println(coins(200))
}

func coins(num int) int {
	arr := []int{100, 50, 20, 10, 5, 2, 1}
	total := 1
	sum := 0
	for i := 0; i <= num/100; i++ {
		for j := 0; j <= num/50; j++ {
			for k := 0; k <= num/20; k++ {
				for l := 0; l <= num/10; l++ {
					for m := 0; m <= num/5; m++ {
						for n := 0; n <= num/2; n++ {
							for o := 0; o <= num; o++ {
								sum = arr[0]*i + arr[1]*j + arr[2]*k + arr[3]*l + arr[4]*m + arr[5]*n + arr[6]*o
								if sum == num {
									total++
								}
							}
						}
					}
				}
			}
		}
	}
	return total
}
