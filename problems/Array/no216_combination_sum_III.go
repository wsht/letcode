package arr

import (
	"fmt"
)

func CombinationSum3(k, n int) [][]int {
	first := make([]int, k+1)
	nc := n
	kc := k - 1
	maxcount := 9 * 10 / 2
	for i := 0; i < k; i++ {
		j := first[i] + 1
		for ; j <= 9; j++ {
			tmp := nc - j
			left := maxcount - (9-kc+1)*(9-kc)/2
			fmt.Println(tmp, left)
			if tmp <= left {
				nc = tmp
				kc--
				first[i+1] = j
				break
			}
		}
	}
	first = first[1:len(first)]
	fmt.Println(first)
	for i := 0; i < k; i++ {

	}
	return [][]int{}
}

func CombinationSum3_2(k, n int) [][]int {
	// path :=
	result := [][]int{}
	BackTrack(&result, []int{}, 1, k, n)

	return result
}

func BackTrack(ret *[][]int, path []int, start, k, target int) {
	if k < 0 || target < 0 {
		return
	}
	if target == 0 && k == 0 {
		fmt.Println("append", path)
		//这里需要copy 否则修改path的值，会导致ret的值被改变
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*ret = append(*ret, pathCopy)
		// fmt.Println(ret)
		return
	}
	for i := start; i <= 9; i++ {
		path = append(path, i)
		// fmt.Println(path)
		BackTrack(ret, path, i+1, k-1, target-i)
		// fmt.Println(ret)
		path = path[:len(path)-1]
	}
	// return ret
}
