package main

import (
	"fmt"
	"strconv"
)

func checksudoku(sudoku [][]int) bool {

	hashmap := make(map[string]int)
	for i := 0; i < len(sudoku); i++ {
		for j := 0; j < len(sudoku[i]); j++ {
			row := strconv.Itoa(sudoku[i][j]) + " row " + strconv.Itoa(i)
			col := strconv.Itoa(sudoku[i][j]) + " col " + strconv.Itoa(j)
			box := strconv.Itoa(sudoku[i][j]) + " box " + strconv.Itoa((i/3)*3) + " " + strconv.Itoa((j/3)*3)

			if _, ok := hashmap[row]; ok {
				fmt.Println(row)
				return false
			}

			if _, ok := hashmap[col]; ok {
				fmt.Println(col)
				return false
			}

			if _, ok := hashmap[box]; ok {
				fmt.Println(box)
				return false
			}

			hashmap[row] = sudoku[i][j]
			hashmap[col] = sudoku[i][j]
			hashmap[box] = sudoku[i][j]

		}
		fmt.Println()
	}

	return true
}

func main() {
	sudoku1 := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	if checksudoku(sudoku1) {
		fmt.Println("Sudoku1 is valid")
	} else {
		fmt.Println("Sudoku1 is not valid")
	}

	sudoku2 := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{2, 3, 4, 5, 6, 7, 8, 9, 1},
		{3, 4, 5, 6, 7, 8, 9, 1, 2},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{5, 6, 7, 8, 9, 1, 2, 3, 4},
		{6, 7, 8, 9, 1, 2, 3, 4, 5},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{8, 9, 1, 2, 3, 4, 5, 6, 7},
		{9, 1, 2, 3, 4, 5, 6, 7, 8},
	}

	if checksudoku(sudoku2) {
		fmt.Println("Sudoku2 is valid")
	} else {
		fmt.Println("Sudoku2 is not valid")
	}

}
