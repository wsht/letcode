package main

import (
	"math"
)

func myAtoi(s string) int {
	sign := 1
	base := 0
	i := 0

	for i < len(s) && string(s[i]) == " " {
		i++
	}
	//handle additional character
	if i < len(s) && (string(s[i]) == "+" || string(s[i]) == "-") {
		if string(s[i]) == "-" {
			sign = -1
		}
		i++
	}
	askIIByte0 := "0"[0]
	askIIByte9 := "9"[0]

	//handle numbic input
	for i < len(s) && s[i] >= askIIByte0 && s[i] <= askIIByte9 {
		if base > math.MaxInt32/10 || (base == math.MaxInt32/10 && int(s[i]-askIIByte0) > 7) {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		base = 10*base + int((s[i] - askIIByte0))
		i++
	}

	return base * sign
}

// func main() {
// 	// result, err := strconv.Atoi("123456")

// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// } else {
// 	// 	fmt.Println(result)
// 	// }
// 	fmt.Println(myAtoi("2147483648"))
// }
