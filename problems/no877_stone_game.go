package main

import (
	"fmt"
)

func stoneGame(piles []int) bool {
	first, second := 0, 0
	i, j := 0, len(piles)-1
	for i < j {
		if piles[i] > piles[j] {
			first += piles[i]
			i++
		} else {
			first += piles[j]
			j--
		}
		if piles[i] > piles[j] {
			second += piles[i]
			i++
		} else {
			second += piles[j]
			j--
		}
	}
	return first > second
}

func main() {
	fmt.Println(stoneGame([]int{5, 3, 4, 5}))
}
