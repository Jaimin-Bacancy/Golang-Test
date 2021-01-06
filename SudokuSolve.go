package main

import (
	"fmt"
	"strconv"
)

func display(mapofint map[string][]int) {
	fmt.Println("Print Sudoku")
	for i := 0; i < 9; i++ {

		intslice := mapofint[strconv.Itoa(i+1)+"row"]
		for _, val := range intslice {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}

func checkvalidvalue(mapofint map[string][]int, currentrow, currentcol, checkval int) bool {

	//column check value is already exits
	colintslice := mapofint[strconv.Itoa(currentrow+1)+"row"]
	for i := 0; i < 9; i++ {
		if colintslice[i] == checkval {

			return false
		}
	}

	//row check value is already exits
	for i := 0; i < 9; i++ {

		rowintslice := mapofint[strconv.Itoa(i+1)+"row"]
		if rowintslice[currentcol] == checkval {
			return false
		}
	}

	//3x3 box check value is already exits
	startingcol := (currentcol / 3) * 3
	startingrow := (currentrow / 3) * 3
	for i := startingrow; i < startingrow+3; i++ {
		boxintslice := mapofint[strconv.Itoa(i+1)+"row"]
		for j := startingcol; j < startingcol+3; j++ {

			if boxintslice[j] == checkval {

				return false
			}
		}

	}
	return true
}

func sudokuSolve(mapofint map[string][]int, row, col int) {

	//if current row at after end of row sudoku is solve
	if row == 9 {
		display(mapofint)
		return
	}

	nextcol := 0
	nextrow := 0

	//reach at last column if current column is last colum then increment current row by 1 and intialize column with 0.
	//otherwise increment current column and row remain yes it is.
	if col == 8 {
		nextrow = row + 1
		nextcol = 0
	} else {
		nextrow = row
		nextcol = col + 1
	}

	//take slice for current row
	intslice := mapofint[strconv.Itoa(row+1)+"row"]

	//check current column value is zero
	if intslice[col] != 0 {

		//skip the current column value
		sudokuSolve(mapofint, nextrow, nextcol)
	} else {

		for val := 1; val <= 9; val++ {

			//check all possiblities of value 1 to 9
			//checkvalidvalue it check the column , row and box if value is not exist than true
			//assign value of not exitsting
			if checkvalidvalue(mapofint, row, col, val) {
				intslice[col] = val
				sudokuSolve(mapofint, nextrow, nextcol)
				//if in last column last possible value is already exits then backtrack and reset value to 0
				intslice[col] = 0
			}
		}
	}

}

func main() {

	mapofint := map[string][]int{
		"1row": []int{3, 0, 6, 0, 0, 8, 4, 0, 0},
		"2row": []int{5, 2, 0, 0, 0, 0, 0, 0, 0},
		"3row": []int{0, 8, 7, 0, 0, 0, 0, 3, 1},
		"4row": []int{0, 0, 3, 0, 1, 0, 0, 8, 0},
		"5row": []int{9, 0, 0, 8, 6, 3, 0, 0, 5},
		"6row": []int{0, 5, 0, 0, 9, 0, 6, 0, 0},
		"7row": []int{1, 3, 0, 0, 0, 0, 2, 5, 0},
		"8row": []int{0, 0, 0, 0, 0, 0, 0, 7, 4},
		"9row": []int{0, 0, 5, 2, 0, 6, 3, 0, 0},
	}
	sudokuSolve(mapofint, 0, 0)

	mapofint2 := map[string][]int{
		"1row": []int{5, 0, 4, 6, 7, 8, 0, 1, 0},
		"2row": []int{6, 7, 0, 1, 0, 5, 3, 4, 8},
		"3row": []int{0, 9, 0, 3, 4, 2, 5, 0, 7},
		"4row": []int{8, 0, 9, 0, 0, 1, 4, 2, 3},
		"5row": []int{0, 2, 0, 8, 5, 0, 7, 9, 1},
		"6row": []int{7, 1, 3, 9, 0, 4, 8, 0, 6},
		"7row": []int{9, 0, 1, 5, 3, 7, 2, 8, 0},
		"8row": []int{2, 8, 0, 0, 0, 0, 6, 3, 5},
		"9row": []int{3, 0, 5, 2, 8, 6, 0, 7, 9},
	}
	sudokuSolve(mapofint2, 0, 0)
}
