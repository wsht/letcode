package main

import (
	"fmt"
	"math"
)

func intToRoman(num int) string {

	if num == 0 {
		return ""
	}

	if num > 4000 {
		return ""
	}

	numCopy := num
	pow := 0
	for numCopy/10 > 0 {
		pow++
		numCopy /= 10
	}

	result := ""

	max := int(math.Pow10(pow))

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

func switchRoman(max int, value int) string {
	switch max {
	case 1000:
		//the value is less than 3999 so just return M
		char := "M"
		ret := ""
		for i := 0; i < value; i++ {
			ret += char
		}
		return ret
	case 100:
		char := "C"
		if value == 5 {
			return "D"
		}
		ret := ""
		for i := 0; i < value; i++ {
			ret += char
		}
		return ret
	case 10:
		char := "X"
		if value == 5 {
			return "L"
		}
		ret := ""
		for i := 0; i < value; i++ {
			ret += char
		}
		return ret
	case 1:
		char := "I"
		if value == 5 {
			return "V"
		}
		ret := ""
		for i := 0; i < value; i++ {
			ret += char
		}
		return ret
	}

	return ""
}

func main() {
	fmt.Println(intToRoman(3999))
}
