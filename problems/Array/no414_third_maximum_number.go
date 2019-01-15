package arr

import (
	"fmt"
)

func ThridMax(nums []int) int {
	topMax := []int{}
	lens := 3
	for i := 0; i < len(nums); i++ {
		topMax = TopMaxHelper(topMax, nums[i], lens)
		fmt.Println(topMax)
	}
	if len(topMax) == 0 {
		return 0
	} else if len(topMax) < lens {
		return topMax[0]
	} else {
		return topMax[lens-1]
	}

}

func TopMaxHelper(top []int, value, lens int) []int {
	for i := 0; i < len(top); i++ {
		if top[i] == value {
			return top
		}
	}
	tmp := value
	dlval := tmp
	fmt.Println("dlvalue", dlval)
	for i := 0; i < len(top); i++ {
		if tmp > top[i] {
			dlval = top[i]
			top[i] = tmp
			tmp = dlval
		}
	}
	if len(top) < lens {
		top = append(top, dlval)
	}

	return top
}
