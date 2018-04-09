package main

import (
	"math"
)

func reverseInteger(x int) int {
	rev := 0
	for x != 0 {
		rev = rev*10 + (x % 10)
		x = x / 10
		if rev > math.MaxInt32 || rev < math.MinInt32 {
			return 0
		}
	}

	return rev
}

// func main() {
// 	fmt.Println(reverseInteger(-120))
// }
