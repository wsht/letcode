package arr

func SearchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

//减少查询次数， 二分查找法
func SearchInsertBinary(nums []int, target int) int {
	low := 0
	hight := len(nums) - 1
	for low <= hight {
		mid := (low + hight) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			hight = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
