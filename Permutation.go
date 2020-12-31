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

	fmt.Println("value of k is ", k)
	for i := 1; i <= n; i++ {

		index := k / fact[n-i]
		if k%fact[n-i] == 0 {
			index -= 1
		}
		fmt.Println("k value is ", k)
		fmt.Println("group size ", fact[n-i])
		fmt.Println("index is ", index)
		result += strconv.Itoa(digit[index])
		fmt.Println("Result is ", result)
		digit = append(digit[:index], digit[index+1:]...)
		fmt.Println(digit)
		k = k - fact[n-i]*index
		fmt.Println("k value is ", k)
		fmt.Println(i, " is completed")

	}

	return result
}

func main() {

	fmt.Println(calculatePermutation(3, 6))
}
