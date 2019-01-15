package backtrack

func Permute(nums []int) [][]int {
	ret := [][]int{}
	list := []int{}
	vmap := map[int]int{}
	PermuteHelper(&ret, nums, list, vmap)
	return ret
}

func PermuteHelper(ret *[][]int, nums, list []int, vmap map[int]int) {
	if len(list) == len(nums) {
		listcp := make([]int, len(list))
		copy(listcp, list)
		*ret = append(*ret, listcp)
		return
	}

	for i := 0; i < len(nums); i++ {
		val, ok := vmap[nums[i]]
		if !ok || val == 0 {
			vmap[nums[i]] = 1
			list = append(list, nums[i])
			PermuteHelper(ret, nums, list, vmap)
			vmap[nums[i]] = 0
			list = list[:len(list)-1]
		}
	}
}
