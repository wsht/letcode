package arr

func NumMagicSquaresInside(grid [][]int) int {
	n := len(grid)
	if n < 3 {
		return 0
	}
	res := 0
	for i := 0; i < n-2; i++ {
		for j := 0; j < n-2; j++ {
			if grid[i+1][j+1] != 5 {
				continue
			}
			if MagicHelper(i, j, grid) {
				res++
			}
		}
	}
	return res
}

func MagicHelper(i, j int, grid [][]int) bool {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if grid[i+x][j+y] > 9 || grid[i+x][j+y] < 1 {
				return false
			}
		}
	}

	for x := 0; x < 3; x++ {
		if grid[i+x][j]+grid[i+x][j+1]+grid[i+x][j+2] != 15 {
			return false
		}
		if grid[i][j+x]+grid[i+1][j+x]+grid[i+2][j+x] != 15 {
			return false
		}
	}
	if grid[i][j]+grid[i+1][j+1]+grid[i+2][j+2] != 15 {
		return false
	}
	if grid[i+2][j]+grid[i+1][j+1]+grid[i][j+2] != 15 {
		return false
	}

	return true
}
