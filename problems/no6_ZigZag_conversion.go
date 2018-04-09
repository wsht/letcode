package main

func zigZagConvert(s string, numRows int) string {
	strLen := len(s)
	cycle := 2*numRows - 2
	result := ""
	if numRows < 2 {
		return s
	}
	for i := 0; i < numRows; i++ {
		for j := i; j < strLen; j = j + cycle {
			result += string(s[j])
			secondJ := cycle - 2*i + j
			if i != 0 && i != numRows-1 && secondJ < strLen {
				result += string(s[secondJ])
			}
		}
	}

	return result
}

// func main() {
// 	fmt.Println(zigZagConvert("PAYPALISHIRING", 2))
// }
