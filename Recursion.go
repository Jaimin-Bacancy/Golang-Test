package main

import "fmt"

func print(count int) {
	fmt.Println("Power Function calls ", count, " time")
}

func power(a, b int) int {
	if b == 0 {
		return 1
	}
	defer print(b)
	return a * power(a, b-1)

}

func main() {
	fmt.Println("!!Recursion code!!")
	var a, b int
	fmt.Println("Enter value of a and b")
	fmt.Scan(&a)
	fmt.Scan(&b)
	c := power(a, b)
	fmt.Println("Power of ", a, "^", b, "is ", c)
}
