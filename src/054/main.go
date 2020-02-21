package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"../helpers"
)

type result struct {
	value   int
	kickers []int
	name    string
}

func main() {
	fmt.Println(pokerHands())
}

func pokerHands() int {
	content, _ := ioutil.ReadFile("./poker.txt")
	hands := string(content)
	handsArr := strings.Split(hands, "\n")
	player1 := [][]string{}
	player2 := [][]string{}
	for i := 0; i < len(handsArr); i++ {
		cards := strings.Split(handsArr[i], " ")
		player1 = append(player1, []string{cards[0], cards[1], cards[2], cards[3], cards[4]})
		player2 = append(player2, []string{cards[5], cards[6], cards[7], cards[8], cards[9]})
	}
	results1 := []result{}
	results2 := []result{}
	for i := 0; i < len(player1); i++ {
		results1 = append(results1, evaluateHand(player1[i]))
		results2 = append(results2, evaluateHand(player2[i]))
	}
	win1 := 0
	win2 := 0
	for i := 0; i < len(results1); i++ {
		if results1[i].value > results2[i].value {
			win1++
		} else if results2[i].value > results1[i].value {
			win2++
		} else if results1[i].value == results2[i].value {
			for j := 0; j < len(results1[i].kickers); j++ {
				if results1[i].kickers[j] > results2[i].kickers[j] {
					win1++
					break
				} else if results2[i].kickers[j] > results1[i].kickers[j] {
					win2++
					break
				}
			}
		}
	}

	return win1
}

func evaluateHand(hand []string) result {
	calcResult := royalFlush(hand)
	if calcResult.value == 9 {
		return calcResult
	}
	calcResult = straightFlush(hand)
	if calcResult.value == 8 {
		return calcResult
	}
	calcResult = fourOfAKind(hand)
	if calcResult.value == 7 {
		return calcResult
	}
	calcResult = fullHouse(hand)
	if calcResult.value == 6 {
		return calcResult
	}
	calcResult = flush(hand)
	if calcResult.value == 5 {
		return calcResult
	}
	calcResult = straight(hand)
	if calcResult.value == 4 {
		return calcResult
	}
	calcResult = trips(hand)
	if calcResult.value == 3 {
		return calcResult
	}
	calcResult = twoPair(hand)
	if calcResult.value == 2 {
		return calcResult
	}
	calcResult = pair(hand)
	if calcResult.value == 1 {
		return calcResult
	}
	return highCard(hand)
}

func royalFlush(hand []string) result {
	colors := "DHSC"
	for i := 0; i < len(colors); i++ {
		if helpers.SliceContainsString(hand, fmt.Sprintf("A%c", colors[i])) &&
			helpers.SliceContainsString(hand, fmt.Sprintf("K%c", colors[i])) &&
			helpers.SliceContainsString(hand, fmt.Sprintf("Q%c", colors[i])) &&
			helpers.SliceContainsString(hand, fmt.Sprintf("J%c", colors[i])) &&
			helpers.SliceContainsString(hand, fmt.Sprintf("T%c", colors[i])) {
			return result{value: 9, kickers: []int{}, name: "Royal Flush"}
		}
	}
	return result{value: 0}
}

func straightFlush(hand []string) result {
	ranks := "23456789TJQKA"
	for i := 0; i < len(hand); i++ {
		split := strings.Split(hand[i], "")
		index := strings.Index(ranks, split[0])
		if index > 3 {
			if helpers.SliceContainsString(hand, fmt.Sprintf("%a%b", ranks[index-1], split[1])) &&
				helpers.SliceContainsString(hand, fmt.Sprintf("%a%b", ranks[index-2], split[1])) &&
				helpers.SliceContainsString(hand, fmt.Sprintf("%a%b", ranks[index-3], split[1])) &&
				helpers.SliceContainsString(hand, fmt.Sprintf("%a%b", ranks[index-4], split[1])) {
				return result{value: 8, kickers: []int{index}, name: fmt.Sprintf("Straight Flush %a", split[0])}
			}
		} else if index == 3 {
			if helpers.SliceContainsString(hand, fmt.Sprintf("%a%b", ranks[index-1], split[1])) &&
				helpers.SliceContainsString(hand, fmt.Sprintf("%a%b", ranks[index-2], split[1])) &&
				helpers.SliceContainsString(hand, fmt.Sprintf("%a%b", ranks[index-3], split[1])) &&
				helpers.SliceContainsString(hand, fmt.Sprintf("A%a", split[1])) {
				return result{value: 8, kickers: []int{index}, name: fmt.Sprintf("Straight Flush %a", split[0])}
			}
		}
	}
	return result{value: 0}
}

func fourOfAKind(hand []string) result {
	ranks := "23456789TJQKA"
	values := ""
	for i := 0; i < len(hand); i++ {
		values += strings.Split(hand[i], "")[0]
	}
	for i := 0; i < len(ranks); i++ {
		count := strings.Count(values, ranks[i:i+1])
		if count == 4 {
			return result{value: 7, kickers: []int{i}}
		}
	}
	return result{value: 0}
}

func fullHouse(hand []string) result {
	ranks := "23456789TJQKA"
	values := ""
	for i := 0; i < len(hand); i++ {
		values += strings.Split(hand[i], "")[0]
	}
	for i := 12; i >= 0; i-- {
		threes := strings.Count(values, ranks[i:i+1])
		if threes == 3 {
			for j := 12; j >= 0; j-- {
				twos := strings.Count(values, ranks[j:j+1])
				if twos == 2 {
					return result{value: 6, kickers: []int{getRank(ranks[i : i+1]), getRank(ranks[j : j+1])}}
				}
			}
		}
	}
	return result{value: 0}
}

func getRank(rank string) int {
	switch rank {
	case "2":
		return 0
	case "3":
		return 1
	case "4":
		return 2
	case "5":
		return 3
	case "6":
		return 4
	case "7":
		return 5
	case "8":
		return 6
	case "9":
		return 7
	case "T":
		return 8
	case "J":
		return 9
	case "Q":
		return 10
	case "K":
		return 11
	case "A":
		return 12
	}

	return 0
}

func flush(hand []string) result {
	ranks := "23456789TJQKA"
	values := ""
	colors := "DHSC"
	kickers := []string{}
	for i := 0; i < len(hand); i++ {
		values += strings.Split(hand[i], "")[1]
		kickers = append(kickers, strings.Split(hand[i], "")[0])
	}
	for i := 0; i < len(colors); i++ {
		count := strings.Count(values, colors[i:i+1])
		if count == 5 {
			resultKickers := []int{}
			for j := 12; j >= 0; j-- {
				if helpers.SliceContainsString(kickers, ranks[j:j+1]) {
					resultKickers = append(resultKickers, j)
				}
			}
			return result{value: 5, kickers: resultKickers}
		}
	}
	return result{value: 0}
}

func straight(hand []string) result {
	ranks := "23456789TJQKA"
	values := ""
	for i := 0; i < len(hand); i++ {
		values += strings.Split(hand[i], "")[0]
	}
	for i := 12; i >= 4; i-- {
		if strings.Index(values, ranks[i:i+1]) != -1 &&
			strings.Index(values, ranks[i-1:i]) != -1 &&
			strings.Index(values, ranks[i-2:i-1]) != -1 &&
			strings.Index(values, ranks[i-3:i-2]) != -1 &&
			strings.Index(values, ranks[i-4:i-3]) != -1 {
			return result{value: 4, kickers: []int{i}}
		}
		if strings.Index(values, ranks[3:4]) != -1 &&
			strings.Index(values, ranks[2:3]) != -1 &&
			strings.Index(values, ranks[1:2]) != -1 &&
			strings.Index(values, ranks[0:1]) != -1 &&
			strings.Index(values, ranks[12:13]) != -1 {
			return result{value: 4, kickers: []int{3}}
		}
	}
	return result{value: 0}
}

func trips(hand []string) result {
	ranks := "23456789TJQKA"
	values := ""
	for i := 0; i < len(hand); i++ {
		values += strings.Split(hand[i], "")[0]
	}
	for i := 12; i >= 0; i-- {
		count := strings.Count(values, ranks[i:i+1])
		if count == 3 {
			kickers := []int{i}
			for j := 12; j >= 0; j-- {
				if strings.Index(values, ranks[j:j+1]) != -1 {
					kickers = append(kickers, j)
				}
			}
			return result{value: 3, kickers: kickers}
		}
	}
	return result{value: 0}
}

func twoPair(hand []string) result {
	ranks := "23456789TJQKA"
	values := ""
	for i := 0; i < len(hand); i++ {
		values += strings.Split(hand[i], "")[0]
	}
	for i := 12; i >= 0; i-- {
		count1 := strings.Count(values, ranks[i:i+1])
		if count1 == 2 {
			for j := i - 1; j >= 0; j-- {
				count2 := strings.Count(values, ranks[j:j+1])
				if count2 == 2 {
					kickers := []int{i, j}
					for k := 12; k >= 0; k-- {
						if strings.Index(values, ranks[k:k+1]) != -1 && k != i && k != j {
							kickers = append(kickers, k)
							return result{value: 2, kickers: kickers}
						}
					}
				}
			}
		}
	}
	return result{value: 0}
}

func pair(hand []string) result {
	ranks := "23456789TJQKA"
	values := ""
	for i := 0; i < len(hand); i++ {
		values += strings.Split(hand[i], "")[0]
	}
	for i := 12; i >= 0; i-- {
		count := strings.Count(values, ranks[i:i+1])
		if count == 2 {
			kickers := []int{i}
			for k := 12; k >= 0; k-- {
				if strings.Index(values, ranks[k:k+1]) != -1 && k != i {
					kickers = append(kickers, k)
					return result{value: 1, kickers: kickers}
				}
			}
		}
	}
	return result{value: 0}
}

func highCard(hand []string) result {
	ranks := "23456789TJQKA"
	values := ""
	for i := 0; i < len(hand); i++ {
		values += strings.Split(hand[i], "")[0]
	}
	kickers := []int{}
	for k := 12; k >= 0; k-- {
		if strings.Index(values, ranks[k:k+1]) != -1 {
			kickers = append(kickers, k)
		}
	}
	return result{value: 0, kickers: kickers}
}
