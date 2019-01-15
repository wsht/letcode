package arr

import (
	"fmt"
	"math"
)

func Subsets(nums []int) [][]int {
	ret := [][]int{}
	// ret = append(ret, []int{})
	cmbo := []int{}
	subsetsHelper(&ret, cmbo, nums, -1)
	return ret
}

//backtrack
func subsetsHelper(ret *[][]int, cmbo []int, nums []int, start int) {
	max := len(nums)
	if len(cmbo) > max {
		return
	}
	cmbocp := make([]int, len(cmbo))
	copy(cmbocp, cmbo)
	fmt.Println(cmbocp)
	*ret = append(*ret, cmbocp)

	for i := start + 1; i < max; i++ {
		val := nums[i]
		cmbo = append(cmbo, val)
		fmt.Println(cmbo)
		subsetsHelper(ret, cmbo, nums, i)
		cmbo = cmbo[:len(cmbo)-1]
	}
}

/**
for example，[1,2,3] -> pow(2,3) -> all bin: 000,001,010,011,....... -> 101 means subset [1,3], 001 means subset[3], -> So, the bit shift part is for knowing if there is a 1, if 1 means num[i] should be appended

I just read a book on competitive programming and the authors says int's binary representations are frequently used to represent all combinations of n elements. Think about each position as an indicator whether the element at this position is "in" or "out". Then we can translate the valid index into a list. Since binary operation is very fast, although this is an exhaustive search we can still get around with the time limit.

Book link : https://www.amazon.com/Competitive-Programming-3rd-Steven-Halim/dp/B00FG8MNN8

https://www.mathsisfun.com/sets/power-set.html

*/
func SubsetsWithBitManipulation(nums []int) [][]int {
	elemnums := len(nums)
	subsetnums := int(math.Pow(float64(2), float64(elemnums)))
	subsets := make([][]int, subsetnums)
	for i := 0; i < elemnums; i++ {
		for j := 0; j < subsetnums; j++ {
			if (j>>uint32(i))&1 == 1 {
				subsets[j] = append(subsets[j], nums[i])
			}
		}
	}

	return subsets
}
