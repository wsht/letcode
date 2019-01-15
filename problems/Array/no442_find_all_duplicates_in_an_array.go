package arr

import (
	"fmt"
	"math"
)

//Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array),

func FindDuplicates(nums []int) []int {
	// ret := []int{}
	n := len(nums)
	//用角标统计每个数字出现的次数
	for i := 0; i < n; i++ {
		tmp := nums[i] - 1
		//已经统计过了，跳过
		if tmp < 0 {
			continue
		}
		tmpValue := nums[tmp]
		if nums[tmp] > 0 {
			nums[tmp] = 0
		}
		nums[tmp]--
		//角标与值不同的时候，防止循环操作
		if i != tmp {
			//目标值没有被统计过
			if tmpValue > 0 {
				nums[i] = tmpValue
			} else {
				nums[i] = 0
			}
			i--
		}
	}
	//统计结果
	ret := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] <= -2 {
			ret = append(ret, i+1)
		}
	}

	return ret
}

func FindDuplicates2(nums []int) []int {
	res := []int{}
	for i := 0; i < len(nums); i++ {
		index := int(math.Abs(float64(nums[i]))) - 1
		if nums[index] < 0 {
			res = append(res, index+1)
		} else {
			nums[index] *= -1
			fmt.Println(nums)
		}
	}
	return res
}
