package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(findOrigins())
}

func findOrigins() int {
	content, _ := ioutil.ReadFile("./p102_triangles.txt")
	triangles := string(content)
	arr := strings.Split(triangles, "\n")
	arr = arr[:len(arr)-1]
	result := 0

	for _, row := range arr {
		coordinates := strings.Split(row, ",")
		if doesContain(coordinates) {
			result++
		}
	}

	return result
}

func doesContain(arr []string) bool {
	Ax, _ := strconv.Atoi(arr[0])
	Ay, _ := strconv.Atoi(arr[1])
	Bx, _ := strconv.Atoi(arr[2])
	By, _ := strconv.Atoi(arr[3])
	Cx, _ := strconv.Atoi(arr[4])
	Cy, _ := strconv.Atoi(arr[5])

	area := float64(Ax*(By-Cy)+Bx*(Cy-Ay)+Cx*(Ay-By)) / 2.0
	a := float64(Bx*Cy+Cx*(-By)) / 2.0
	b := float64(Ax*(-Cy)+Cx*Ay) / 2.0
	c := float64(Ax*By+Bx*(-Ay)) / 2.0

	return math.Abs(area) == math.Abs((a))+math.Abs(b)+math.Abs(c)
}
