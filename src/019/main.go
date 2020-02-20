package main

import "fmt"

func main() {
	fmt.Println(sundays())
}

func sundays() int {
	sunday := 0
	days := -5
	for year := 1901; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			switch month {
			case 1, 3, 5, 7, 8, 10, 12:
				days += 31
			case 4, 6, 9, 11:
				days += 30
			case 2:
				if year%4 == 0 {
					days += 29
				} else {
					days += 28
				}
			}
			if days%7 == 0 {
				if sunday == 0 {
					fmt.Println(month, year)
				}
				sunday++
			}
		}
	}
	return sunday
}
