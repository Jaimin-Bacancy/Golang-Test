package main

import "fmt"

func checkthreecrossthree(sudoku [][]int, m int, n int, curval int) bool {
	//defer fmt.Println("Completed check")
	count := 0
	//fmt.Println(curval, " ", m, " ", n)
	//fmt.Println()
	for i := m; i < m+3; i++ {
		for j := n; j < n+3; j++ {
			//fmt.Print(sudoku[i][j], " ")
			if curval == sudoku[i][j] {
				count++
			}
		}
		//fmt.Println()
	}
	if count > 1 {
		return false
	}
	return true
}

func checksudoku(sudoku [][]int) bool {

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			curval := sudoku[i][j]

			for k := j + 1; k < 9; k++ {
				if curval == sudoku[i][k] || curval < 1 || curval > 9 {
					return false
				}
			}

			for k := i + 1; k < 9; k++ {
				if curval == sudoku[k][j] || curval < 1 || curval > 9 {
					return false
				}
			}

			m := (i / 3) * 3
			n := (j / 3) * 3

			if !(checkthreecrossthree(sudoku, m, n, curval)) {
				return false
			}

		}
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
