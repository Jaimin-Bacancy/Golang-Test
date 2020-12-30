package main

import (
	"fmt"
	"strconv"
)

func calculatePermutation(n int, k int) string {
	fact := make([]int, n+1)
	digit := make([]int, n)
	var result string = ""
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = i * fact[i-1]
		digit[i-1] = i
	}
	fmt.Println(fact)
	fmt.Println(digit)
	k--
	for i := 1; i <= n; i++ {

		index := k / fact[n-i]
		result += strconv.Itoa(digit[index])
		digit = append(digit[:index], digit[index+1:]...)
		fmt.Println(digit)
		k = k - fact[n-i]*index
	}

	return result
}

func main() {
	fmt.Println(calculatePermutation(3, 3))
}
