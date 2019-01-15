package backtrack

func CountArrangement(N int) int {
	total := 0
	visited := make([]bool, N+1)
	CountArrangementHelper(N, visited, 1, &total)
	return total
}

func CountArrangementHelper(N int, visited []bool, pos int, total *int) {
	if pos > N {
		*total += 1
		return
	}
	for i := 1; i <= N; i++ {
		if !visited[i] && (pos%i == 0 || i%pos == 0) {
			// fmt.Println(pos, i)
			visited[i] = true
			CountArrangementHelper(N, visited, pos+1, total)
			visited[i] = false
		}
	}
}

func countArrangement(N int) int {
	if N <= 0 {
		return 0
	}
	nums := make([]int, N)
	for i := 0; i < N; i++ {
		nums[i] = i + 1

	}
	return count(nums, N)
}

func count(nums []int, n int) int {
	if n == 0 {
		return 1
	}
	result := 0
	for i := 0; i < n; i++ {
		if n%nums[i] == 0 || nums[i]%n == 0 {
			nums[i], nums[n-1] = nums[n-1], nums[i]
			result += count(nums, n-1)
			nums[i], nums[n-1] = nums[n-1], nums[i]
		}
	}
	return result
}
