package main

import "fmt"

func transpose(matA [][]int) {

	var matRes [3][3]int

	for i := 0; i < len(matA); i++ {
		for j := 0; j < len(matA[i]); j++ {
			matRes[i][j] = matA[j][i]
		}
	}

	fmt.Println(matRes)
}

func main() {
	fmt.Println("Transpose of Matrix")
	matA := [][]int{{9, 7, 5}, {4, 9, 2}, {9, 7, 5}}
	transpose(matA)
}
