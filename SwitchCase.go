package main

import (
	"fmt"
	"math"
)

func case1(binary string) {
	var countzero float64
	var res float64 = 0
	for i := 0; i < len(binary); i++ {
		digit := string([]rune(binary)[i])
		if digit == "1" {
			countzero = 0
		} else {
			countzero++
			res = math.Max(res, countzero)
		}
	}
	nres := int64(res)
	//fmt.Println(nres)
	if nres >= 3 {
		fmt.Println("000-ODD")
	} else {
		fmt.Println(">3 0s-ODD")
	}

}

func case2(binary string) {
	var countone float64
	var res float64 = 0
	for i := 0; i < len(binary); i++ {
		digit := string([]rune(binary)[i])
		if digit == "0" {
			countone = 0
		} else {
			countone++
			res = math.Max(res, countone)
		}
	}
	nres := int64(res)
	//fmt.Println(nres)
	if nres <= 5 {
		fmt.Println(">5 1s-EVEN")
	}
}

func main() {
	var binary string

	fmt.Println("!!Switch Case!!")
	fmt.Println("Enter Binary Number")
	fmt.Scan(&binary)

	switch number := binary; string([]rune(binary)[len(binary)-1]) {
	//case 1 : check 3 consecutive zero or more and binary number is odd
	//print('000-odd')

	//case 1 : check 3 consecutive zero lessthan it and binary number is odd
	//print('>3 0s-ODD')

	case "1":
		//fmt.Println("odd number")
		case1(number)
	case "0":
		//fmt.Println("Even number")
		case2(number)
	default:
		fmt.Println("Default is execute")
	}
}
