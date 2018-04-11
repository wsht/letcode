package main

func romanToInteger(roman string) int {
	romanMap := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}

	strLen := len(roman)

	sum := 0
	for i := 0; i < strLen-1; i++ {
		curValue := romanMap[string(roman[i])]
		nextValue := romanMap[string(roman[i+1])]
		if curValue >= nextValue {
			sum += curValue
		} else {
			sum -= curValue
		}
	}

	sum += romanMap[string(roman[strLen-1])]
	return sum
}

// func main() {
// 	fmt.Println(romanToInteger("XXX"))
// }
