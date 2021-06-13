package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type SudokuSquare struct {
	value   int
	options []int
}

func main() {
	fmt.Println(findSudokuSum())
}

func findSudokuSum() int {
	content, _ := ioutil.ReadFile("./p096_sudoku.txt")
	sudoku := strings.Split(string(content), "\n")

	sum := 0

	for i := 0; i < len(sudoku); i += 10 {
		num := solveSudoku(sudoku[i : i+10])

		sum += num
	}

	return sum
}

func solveSudoku(arr []string) int {
	sudoku := [9][9]SudokuSquare{}

	for i := 1; i < len(arr); i++ {
		row := parseRow(arr[i])
		sudoku[i-1] = row
	}

	sudoku = fillOptions(sudoku)

	sudoku = checkSingleOption(sudoku)

	if checkSolve(sudoku) {
		return parseResult(sudoku)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku[i][j].value == 0 {
				for _, val := range sudoku[i][j].options {
					result := oneSubTest(sudoku, i, j, val)
					if checkSolve(result) {
						return parseResult(result)
					} else if isSudokuValid(result) {
						for i := 0; i < 9; i++ {
							for j := 0; j < 9; j++ {
								if result[i][j].value == 0 {
									for _, val := range result[i][j].options {
										result2 := oneSubTest(result, i, j, val)
										if checkSolve(result2) {
											return parseResult(result2)
										} else if isSudokuValid(result2) {
											for i := 0; i < 9; i++ {
												for j := 0; j < 9; j++ {
													if result2[i][j].value == 0 {
														for _, val := range result2[i][j].options {
															result3 := oneSubTest(result2, i, j, val)
															if checkSolve(result3) {
																return parseResult(result3)
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return 0
}

func oneSubTest(sudoku [9][9]SudokuSquare, i int, j int, val int) [9][9]SudokuSquare {
	sudoku[i][j].value = val
	sudoku[i][j].options = []int{}

	return removeOptions(sudoku, i, j)
}

func isSudokuValid(sudoku [9][9]SudokuSquare) bool {
	for i := 0; i < 9; i++ {
		if testDuplicates(collectRow(sudoku, i)) {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if testDuplicates(collectColumn(sudoku, i)) {
			return false
		}
	}
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if testDuplicates(collectSquare(sudoku, i, j)) {
				return false
			}
		}
	}

	return true
}

func collectRow(sudoku [9][9]SudokuSquare, r int) [9]int {
	row := [9]int{}

	for i := 0; i < 9; i++ {
		row[i] = sudoku[r][i].value
	}

	return row
}

func testDuplicates(arr [9]int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			for j := i + 1; j < len(arr); j++ {
				if arr[i] == arr[j] {
					return true
				}
			}
		}
	}

	return false
}

func collectSquare(sudoku [9][9]SudokuSquare, row int, col int) [9]int {
	square := [9]int{}
	index := 0

	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			square[index] = sudoku[i][j].value
			index++
		}
	}

	return square
}

func collectColumn(sudoku [9][9]SudokuSquare, col int) [9]int {
	column := [9]int{}

	for i := 0; i < 9; i++ {
		column[i] = sudoku[i][col].value
	}

	return column
}

func parseResult(sudoku [9][9]SudokuSquare) int {
	str := strconv.Itoa(sudoku[0][0].value) + strconv.Itoa(sudoku[0][1].value) + strconv.Itoa(sudoku[0][2].value)

	num, _ := strconv.Atoi(str)

	return num
}

func checkSolve(sudoku [9][9]SudokuSquare) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku[i][j].value == 0 {
				return false
			}
		}
	}

	return isSudokuValid(sudoku)
}

func printSudoku(sudoku [9][9]SudokuSquare) {
	for _, val := range sudoku {
		fmt.Println(val)
	}
}

func checkSingleOption(sudoku [9][9]SudokuSquare) [9][9]SudokuSquare {
	for i := 0; i < len(sudoku); i++ {
		for j := 0; j < len(sudoku[i]); j++ {
			if sudoku[i][j].value == 0 && len(sudoku[i][j].options) == 1 {
				sudoku[i][j].value = sudoku[i][j].options[0]
				sudoku[i][j].options = []int{}
				sudoku = removeOptions(sudoku, i, j)
			}
		}
	}

	return sudoku
}

func removeOptions(sudoku [9][9]SudokuSquare, row int, col int) [9][9]SudokuSquare {
	for j := 0; j < len(sudoku[row]); j++ {
		sudoku[row][j].options = removeItem(sudoku[row][j].options, sudoku[row][col].value)
		if len(sudoku[row][j].options) == 1 {
			sudoku[row][j].value = sudoku[row][j].options[0]
			sudoku[row][j].options = []int{}
			sudoku = removeOptions(sudoku, row, j)
		}
	}
	for i := 0; i < len(sudoku); i++ {
		sudoku[i][col].options = removeItem(sudoku[i][col].options, sudoku[row][col].value)
		if len(sudoku[i][col].options) == 1 {
			sudoku[i][col].value = sudoku[i][col].options[0]
			sudoku[i][col].options = []int{}
			sudoku = removeOptions(sudoku, i, col)
		}
	}

	slug := strconv.Itoa(row) + "_" + strconv.Itoa(col)

	switch slug {
	case "0_0", "0_1", "0_2", "1_0", "1_1", "1_2", "2_0", "2_1", "2_2":
		sudoku = removeSquareOptions(sudoku, 0, 0, sudoku[row][col].value)
	case "0_3", "0_4", "0_5", "1_3", "1_4", "1_5", "2_3", "2_4", "2_5":
		sudoku = removeSquareOptions(sudoku, 0, 3, sudoku[row][col].value)
	case "0_6", "0_7", "0_8", "1_6", "1_7", "1_8", "2_6", "2_7", "2_8":
		sudoku = removeSquareOptions(sudoku, 0, 6, sudoku[row][col].value)
	case "3_0", "3_1", "3_2", "4_0", "4_1", "4_2", "5_0", "5_1", "5_2":
		sudoku = removeSquareOptions(sudoku, 3, 0, sudoku[row][col].value)
	case "3_3", "3_4", "3_5", "4_3", "4_4", "4_5", "5_3", "5_4", "5_5":
		sudoku = removeSquareOptions(sudoku, 3, 3, sudoku[row][col].value)
	case "3_6", "3_7", "3_8", "4_6", "4_7", "4_8", "5_6", "5_7", "5_8":
		sudoku = removeSquareOptions(sudoku, 3, 6, sudoku[row][col].value)
	case "6_0", "6_1", "6_2", "7_0", "7_1", "7_2", "8_0", "8_1", "8_2":
		sudoku = removeSquareOptions(sudoku, 6, 0, sudoku[row][col].value)
	case "6_3", "6_4", "6_5", "7_3", "7_4", "7_5", "8_3", "8_4", "8_5":
		sudoku = removeSquareOptions(sudoku, 6, 3, sudoku[row][col].value)
	case "6_6", "6_7", "6_8", "7_6", "7_7", "7_8", "8_6", "8_7", "8_8":
		sudoku = removeSquareOptions(sudoku, 6, 6, sudoku[row][col].value)
	}

	return sudoku
}

func removeSquareOptions(sudoku [9][9]SudokuSquare, row int, col int, val int) [9][9]SudokuSquare {
	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			sudoku[i][j].options = removeItem(sudoku[i][j].options, val)
			if len(sudoku[i][j].options) == 1 {
				sudoku[i][j].value = sudoku[i][j].options[0]
				sudoku[i][j].options = []int{}
				sudoku = removeOptions(sudoku, i, j)
			}
		}
	}

	return sudoku
}

func fillOptions(sudoku [9][9]SudokuSquare) [9][9]SudokuSquare {
	for i := 0; i < len(sudoku); i++ {
		for j := 0; j < len(sudoku[i]); j++ {
			if sudoku[i][j].value == 0 {
				possibleOptions := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
				row := getRowOptions(sudoku, i)
				column := getColumnOptions(sudoku, j)
				square := getSquareOptions(sudoku, i, j)
				for _, val := range row {
					if possibleOptions[val-1] == val {
						possibleOptions[val-1] = 0
					}
				}
				for _, val := range column {
					if possibleOptions[val-1] == val {
						possibleOptions[val-1] = 0
					}
				}
				for _, val := range square {
					if possibleOptions[val-1] == val {
						possibleOptions[val-1] = 0
					}
				}
				for _, val := range possibleOptions {
					if val != 0 {
						sudoku[i][j].options = append(sudoku[i][j].options, val)
					}
				}
			}
		}
	}

	return sudoku
}

func removeItem(arr []int, item int) []int {
	newArr := []int{}

	for _, val := range arr {
		if val != item {
			newArr = append(newArr, val)
		}
	}

	return newArr
}

func getSquareOptions(sudoku [9][9]SudokuSquare, i int, j int) []int {
	square := []int{}
	slug := strconv.Itoa(i) + "_" + strconv.Itoa(j)

	switch slug {
	case "0_0", "0_1", "0_2", "1_0", "1_1", "1_2", "2_0", "2_1", "2_2":
		return getSquare(sudoku, 0, 0)
	case "0_3", "0_4", "0_5", "1_3", "1_4", "1_5", "2_3", "2_4", "2_5":
		return getSquare(sudoku, 0, 3)
	case "0_6", "0_7", "0_8", "1_6", "1_7", "1_8", "2_6", "2_7", "2_8":
		return getSquare(sudoku, 0, 6)
	case "3_0", "3_1", "3_2", "4_0", "4_1", "4_2", "5_0", "5_1", "5_2":
		return getSquare(sudoku, 3, 0)
	case "3_3", "3_4", "3_5", "4_3", "4_4", "4_5", "5_3", "5_4", "5_5":
		return getSquare(sudoku, 3, 3)
	case "3_6", "3_7", "3_8", "4_6", "4_7", "4_8", "5_6", "5_7", "5_8":
		return getSquare(sudoku, 3, 6)
	case "6_0", "6_1", "6_2", "7_0", "7_1", "7_2", "8_0", "8_1", "8_2":
		return getSquare(sudoku, 6, 0)
	case "6_3", "6_4", "6_5", "7_3", "7_4", "7_5", "8_3", "8_4", "8_5":
		return getSquare(sudoku, 6, 3)
	case "6_6", "6_7", "6_8", "7_6", "7_7", "7_8", "8_6", "8_7", "8_8":
		return getSquare(sudoku, 6, 6)
	}

	return square
}

func getSquare(sudoku [9][9]SudokuSquare, row int, col int) []int {
	square := []int{}

	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			if sudoku[i][j].value != 0 {
				square = append(square, sudoku[i][j].value)
			}
		}
	}

	return square
}

func getColumnOptions(sudoku [9][9]SudokuSquare, j int) []int {
	column := []int{}

	for i := 0; i < len(sudoku); i++ {
		if sudoku[i][j].value != 0 {
			column = append(column, sudoku[i][j].value)
		}
	}

	return column
}

func getRowOptions(sudoku [9][9]SudokuSquare, i int) []int {
	row := []int{}

	for j := 0; j < len(sudoku[i]); j++ {
		if sudoku[i][j].value != 0 {
			row = append(row, sudoku[i][j].value)
		}
	}

	return row
}

func parseRow(row string) [9]SudokuSquare {
	parsed := [9]SudokuSquare{}

	for index, char := range row {
		num, _ := strconv.Atoi(string(char))
		parsed[index] = SudokuSquare{value: num, options: []int{}}
	}

	return parsed
}
