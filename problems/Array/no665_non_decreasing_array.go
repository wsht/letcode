package arr

import (
	"fmt"
)

func CheckPossibility(nums []int) bool {
	count := 0
	prev := nums[0]
	n := len(nums)
	if n <= 2 {
		return true
	}

	for i := 1; i < len(nums); i++ {
		fmt.Println(prev, nums[i], i)
		if prev > nums[i] {
			count++
			//这里跑偏了
			if i == 1 {
				prev = nums[i]
			} else {
				prev = nums[i-2]
			}

		} else {
			prev = nums[i]
		}
		if count > 1 {
			return false
		}
	}

	return true
}

func CheckPossibility2(nums []int) bool {
	count := 0
	n := len(nums)
	index := 0
	if n <= 2 {
		return true
	}
	//首先 需要找到非递增点，并且计算个数
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			count++
			index = i - 1
		}
		if count > 1 {
			return false
		}
	}
	//纯递增
	if count == 0 {
		return true
	}
	// fmt.Println(index)
	res := false
	//去除其中一个点，看剩下的是否连续递增
	for i := 0; i <= 1; i++ {
		del := index + i
		prev := 0
		isNot := false
		for j := 0; j < n-1; j++ {
			if prev == del {
				prev++
			}
			if j+1 == del {
				continue
			}
			fmt.Println(i, j, del, prev, j+1)
			if nums[prev] > nums[j+1] {
				isNot = true
				break
			}
			prev++
		}
		if isNot {
			res = res || false
		} else {
			res = true
		}
	}

	return res
}

func CheckPossibility3(nums []int) bool {
	p := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[i+1] {
			if p != -1 {
				return false
			}
			p = i
		}
	}

	return p == -1 || p == 0 || p == len(nums)-2 || nums[p-1] <= nums[p+1] || nums[p] <= nums[p+2]
}
