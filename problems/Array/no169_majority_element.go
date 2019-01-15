package arr

func MajorityElementwithMap(nums []int) int {
	mMap := map[int]int{}
	majority := nums[0]
	for i := 0; i < len(nums); i++ {
		mMap[nums[i]] += 1
		if mMap[nums[i]] > mMap[majority] {
			majority = nums[i]
		}
	}
	return majority
}

//寻找最大的值，如果不等于的话 则进行减法，
//因为前提条件是结果一定存在，并且其值大于等于 n/2
func MajorityElement(nums []int) int {
	count := 0
	candidate := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}
		if nums[i] == candidate {
			count += 1
		} else {
			count -= 1
		}
	}
	return candidate
}
