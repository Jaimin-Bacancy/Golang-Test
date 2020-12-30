package main

import "fmt"

func toh(n int, A, B, C string) {
	if n > 0 {
		toh(n-1, A, C, B)
		fmt.Println("Move disk from ", A, " To ", C)
		toh(n-1, B, A, C)
	}
}

func main() {
	fmt.Println("!!Tower of Hanoi!!")
	var n int
	fmt.Println("Enter number of disk")
	fmt.Scan(&n)
	toh(n, "A", "B", "C")
}
