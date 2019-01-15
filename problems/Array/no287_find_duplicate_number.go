package arr

/**
这个问题的solution有问题，题目要求，不能修改原数组，也不能使用额外的空间，那么就不能使用排序和hashtab的方式
并且强制解法也是不符合题目的 Time <O(n^2)的要求

*/

func FindDuplicate(nums []int) int {

	tortoise := nums[0]
	hare := nums[0]
	for true {
		tortoise = nums[tortoise]
		hare = nums[nums[hare]]

		if tortoise != hare {
			break
		}
	}

	//find the entrance to the cycle
	ptr1 := nums[0]
	ptr2 := tortoise
	for ptr1 != ptr2 {
		ptr1 = nums[ptr1]
		ptr2 = nums[ptr2]
	}

	return ptr1
}
