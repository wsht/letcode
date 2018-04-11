package main

import (
	"math"
)

func intToRoman(num int) string {
	if num == 0 || num > 4000 {
		return ""
	}

	pow := integerLens(num) - 1
	max := int(math.Pow10(pow))
	result := ""
	for max >= 1 {
		value := num / max
		num = num % max
		// fmt.Println(max, num, value)
		if value%5 == 4 {
			result += switchRoman(max, 1) + switchRoman(max*int(math.Pow10(value/5)), (value+1)/10+(value+1)%10)
		} else {
			result += switchRoman(max, (value/5)*5) + switchRoman(max, value%5)
		}
		max /= 10
	}

	return result
}

func integerLens(num int) int {
	iLen := 1
	for num/10 > 0 {
		iLen++
		num /= 10
	}
	return iLen
}

func switchRoman(max int, value int) string {
	defaultChar := ""
	switch max {
	case 1000:
		//the value is less than 3999 so just return M
		defaultChar = "M"
	case 100:
		if value == 5 {
			return "D"
		}
		defaultChar = "C"
	case 10:
		if value == 5 {
			return "L"
		}
		defaultChar = "X"
	case 1:
		if value == 5 {
			return "V"
		}
		defaultChar = "I"
	}
	result := ""
	for i := 0; i < value; i++ {
		result += defaultChar
	}

	return result
}

// func main() {
// 	fmt.Println(intToRoman(3999))
// }
