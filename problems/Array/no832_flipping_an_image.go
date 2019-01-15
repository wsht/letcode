package arr

func FlipAndInvertImage(A [][]int) [][]int {
	ALen := len(A)
	for i := 0; i < ALen; i++ {
		ARowLen := len(A[i])
		//奇数和偶数的区别
		//循环1/2次，分别互换并且取反[i][j] [i][len-j-1]
		//因为[i][j]率先被赋值，所以需要缓存下来
		for j := 0; j < ARowLen/2+ARowLen%2; j++ {
			tmp := A[i][j]
			A[i][j] = 1 ^ A[i][ARowLen-j-1]
			A[i][ARowLen-j-1] = 1 ^ tmp
		}
	}
	return A
}
