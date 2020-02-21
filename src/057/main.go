package main

import (
	"fmt"

	"../helpers"
)

func main() {
	fmt.Println(squareRoots())
}

func squareRoots() int {
	nom := []string{"1", "3"}
	dem := []string{"1", "2"}

	prevNom := "1"
	currentNom := "3"
	resultNom := ""
	prevDem := "1"
	currentDem := "2"
	resultDem := ""
	for i := 0; i < 1000; i++ {
		resultNom = helpers.AddBigNum(helpers.MultiplyBig(currentNom, "2"), prevNom)
		prevNom = currentNom
		currentNom = resultNom
		nom = append(nom, resultNom)
		resultDem = helpers.AddBigNum(helpers.MultiplyBig(currentDem, "2"), prevDem)
		prevDem = currentDem
		currentDem = resultDem
		dem = append(dem, resultDem)
	}

	result := 0
	for i := 0; i < len(nom); i++ {
		if len(nom[i]) > len(dem[i]) {
			result++
		}
	}
	return result
}
