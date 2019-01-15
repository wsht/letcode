package arr

//方法1，排序，产看 nums[n-1] >= 2*nums[n-2]

func DominantIndex(nums []int) int {
	maxMin, maxMax := 0, 0
	maxIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= maxMax {
			maxMin = maxMax
			maxMax = nums[i]
			maxIndex = i
		} else if nums[i] > maxMin {
			maxMin = nums[i]
		}
	}
	if maxMax >= maxMin*2 {
		return maxIndex
	} else {
		return -1
	}
}
