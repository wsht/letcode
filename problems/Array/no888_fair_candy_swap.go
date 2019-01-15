package arr

func FairCandySwap(A []int, B []int) []int {
	sumA := 0
	sumB := 0
	for i := 0; i < len(A); i++ {
		sumA += A[i]
	}

	for j := 0; j < len(B); j++ {
		sumB += B[j]
	}

	between := (sumA + sumB) / 2
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if sumA-A[i]+B[j] == between && sumB-B[j]+A[i] == between {
				return []int{A[i], B[j]}
			}
		}
	}
	return []int{0, 0}
}

//O(n) space O(B)
func FairCandySwapLiner(A, B []int) []int {
	sa, sb := 0, 0
	for i := 0; i < len(A); i++ {
		sa += A[i]
	}
	bmap := map[int]int{}
	for j := 0; j < len(B); j++ {
		sb += B[j]
		bmap[B[j]] = 1
	}
	between := (sb - sa) / 2
	for i := 0; i < len(A); i++ {
		tmp := A[i] + between
		_, ok := bmap[tmp]
		if ok {
			return []int{A[i], tmp}
		}
	}
	return []int{0, 0}
}
