package arr

func ValidMountainArray(A []int) bool {
	res := false
	n := len(A)
	if n < 3 {
		return res
	}
	for i := 1; i < n-1; i++ {
		if i == 1 && A[i-1] > A[i] {
			return false
		}

		//查找峰值，看是否有第二个峰值
		if A[i] > A[i-1] && A[i] > A[i+1] {
			if res {
				return false
			}
			res = true
			//查看又侧是否一直递减
		} else if res && A[i] <= A[i+1] {
			return false
		}
	}
	return res
}

func ValidMountainArrayOnPass(A []int) bool {
	n := len(A)
	i := 0
	for i+1 < n && A[i] < A[i+1] {
		i++
	}

	if i == 0 || i == n-1 {
		return false
	}

	for i+1 < n && A[i] > A[i+1] {
		i++
	}
	return i == n-1
}
