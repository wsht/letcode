package arr

func PivotIndex(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			if nums[len(nums)-1]-nums[i] == 0 {
				return i
			}
		} else if i == len(nums)-1 {
			if nums[len(nums)-2] == 0 {
				return i
			}
		} else if nums[i-1] == nums[len(nums)-1]-nums[i] {
			return i
		}
	}
	return -1
}

func PivotIndexWioutChangeArray(nums []int) int {
	sum := 0
	leftSum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		if leftSum == sum-leftSum-nums[i] {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}
