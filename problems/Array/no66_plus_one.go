package arr

func PlusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	newNumber := make([]int, n+1)
	newNumber[0] = 1
	return newNumber
}
