package arr

import (
	"fmt"
	"math"
)

//这个方法是不对的，没有考虑结果集是否包含 k个结果，
func ConstructArray(n, k int) []int {
	mark := make([]bool, n)
	ret := []int{1}
	mark[0] = true
	between := k
	firstK := k
	for i := 1; i < n; i++ {
		test := []int{int(math.Abs(float64(ret[i-1] + between))), int(math.Abs(float64(ret[i-1] - between)))}
		// fmt.Println(ret, test)
		match := false
		for j := 0; j < 2; j++ {
			if test[j] > 1 && test[j] <= n && !mark[test[j]-1] {
				ret = append(ret, test[j])
				mark[test[j]-1] = true
				match = true
				break
			}
		}
		between--
		if !match {
			i--
			if between == 0 || len(ret) == n {
				if len(ret) < n {
					for true {
						curlen := len(ret)
						if curlen < 2 {
							firstK--
							between = firstK
							mark = make([]bool, n)
							mark[0] = true
							ret = []int{1}
							if firstK <= 0 {
								return ret
							}
							break
						}
						tmpbetween := int(math.Abs(float64(ret[curlen-1] - ret[curlen-2])))
						test := []int{int(math.Abs(float64(ret[curlen-2] + tmpbetween))), int(math.Abs(float64(ret[curlen-2] - tmpbetween)))}
						// fmt.Println(curlen, tmpbetween, test, ret)
						if test[0] == ret[curlen-1] {
							// fmt.Println("ret[curlen-1]==test[0]", ret, test)
							// fmt.Println(test[1], mark)
							// fmt.Println(test[1] > 1, test[1] <= n, !mark[test[1]-1])
							if test[1] > 1 && test[1] <= n && !mark[test[1]-1] {
								// fmt.Println("over here")
								mark[ret[curlen-1]-1] = false
								ret[curlen-1] = test[1]
								mark[ret[curlen-1]-1] = true
								between = tmpbetween - 1
								break
							}
						}
						//如果当前是第二个值,或者第二个值不符合要求 说明这两个取值都不正常，需要向上回退
						mark[ret[curlen-1]-1] = false
						ret = ret[0 : curlen-1]
						i--
					}
				}
			}
		}
		if between == 0 {
			between = k
		}
	}
	return ret
}

//那么 我们可不可以改变一种思路，第一遍只找 K 个不同组合，n-k的时候，使其间隔全为1
func ConstructArrayOnlyK(n, k int) []int {
	ret := make([]int, n)
	max := k + 1
	min := 1
	for i := 0; i <= k; i++ {
		if i%2 == 0 {
			ret[i] = min
			min++
		} else {
			ret[i] = max
			max--
		}
	}
	// fmt.Println(ret)
	max = k + 2
	for i := k + 1; i < n; i++ {
		ret[i] = max
		max++
	}
	// fmt.Println(ret)
	return ret
}

func ConstructArrayResultCheck(ret []int, n, k int) {
	between := make([]int, n)
	numberCheck := make([]int, n)

	for i := 1; i < len(ret); i++ {
		bi := int(math.Abs(float64(ret[i] - ret[i-1])))
		between[bi]++
		numberCheck[i]++
	}
	btc := 0
	for i := 0; i < n; i++ {
		if between[i] >= 1 {
			btc++
		}
		if numberCheck[i] == 0 || numberCheck[i] > 1 {
			fmt.Print(i)
		}
	}
	fmt.Println("")
	fmt.Println(btc, n)

}
