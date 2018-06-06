package main

/**
Given a digit string, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below.

Input:Digit string "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
*/

func letterCombinations(digits string) []string {
	if len(digits) <= 0 {
		return []string{}
	}
	strMapping := []string{"0", "1", "abc", "def", "ghi", "jkl", "mon", "pqrs", "tuv", "wxyz"}
	res := []string{""}

	for i := 0; i < len(digits); i++ {
		tempres := []string{}

		char := strMapping[digits[i]-"0"[0]]
		for c := 0; c < len(char); c++ {
			for j := 0; j < len(res); j++ {
				tempres = append(tempres, res[j]+string(char[c]))
			}
		}

		res = tempres
	}

	return res
}
