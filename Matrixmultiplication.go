package main

import "fmt"

func multiply(matA, matB [][]int) {
	var matRes [3][3]int
	colA := 2

	for i := 0; i < len(matRes); i++ {
		for j := 0; j < len(matRes[i]); j++ {
			for k := 0; k < colA; k++ {
				matRes[i][j] += matA[i][k] * matB[k][j]
			}
		}
	}
	fmt.Println(matRes)

}

func main() {
	fmt.Println("Matrix Multiplication 3x2 and 2x3 -> 3x3")
	matA := [][]int{{3, 6}, {5, 7}, {5, 2}}
	matB := [][]int{{9, 7, 5}, {4, 9, 2}}
	multiply(matA, matB)

}
