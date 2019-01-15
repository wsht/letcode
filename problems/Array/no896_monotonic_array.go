package arr

func IsMonotonic(A []int) bool {
	isAdd := A[len(A)-1]-A[0] > 0
	for i := 1; i < len(A); i++ {
		res := A[i] - A[i-1]
		if res != 0 && res > 0 != isAdd {
			return false
		}
	}
	return true
}
