package main

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reverseNumber := 0
	for x > reverseNumber {
		reverseNumber = reverseNumber*10 + (x % 10)
		x = x / 10
	}
	if x == reverseNumber || x == reverseNumber/10 {
		return true
	}
	return false
}

// func main() {
// 	fmt.Println(isPalindrome(10001))
// }
