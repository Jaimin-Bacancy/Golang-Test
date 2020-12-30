package main

import "fmt"

func jumpgame(jumps []int) int {
	if len(jumps) == 1 || jumps[0] == 0 {
		return 0
	}

	minjump := 1
	farjump := jumps[0]
	cur := jumps[0]

	for i := 1; i < len(jumps); i++ {
		if i == len(jumps)-1 {
			return minjump
		}

		if i+jumps[i] >= farjump {
			farjump = i + jumps[i]
		}

		if i == cur {
			minjump++
			cur = farjump
		}
	}
	return minjump

}

func main() {
	fmt.Println("JUMP GAME")
	jumps := []int{2, 3, 1, 1, 4}
	fmt.Println("Minimum Jump -> ", jumpgame(jumps))
}
