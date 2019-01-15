package arr

import (
	"math"
)

/**
For this question, it does not matter which you choose. Since we are computing the area of the island, we MUST traverse EVERY SINGLE element in the island; whether we use BFS or DFS, it does not matter -- the time complexity will be the same.

BFS, however, is sometimes preferable. Suppose you had a 1000x1000 grid, and you wanted to find the distance between two points. In this case, you DON'T need to exhaust the entire one million elements (or rather, you only do in the worst case). Consider that your two points happened to be only 5 elements apart. In that case, BFS would find the two pretty quickly, since BFS stays close to the original point. With DFS, however, you'd be flying down the wrong path for a great distance, and it could wind up taking very long.

Even better, think about something like Facebook -- what if you wanted to determine whether two people are connected through a chain of friends? If someone is only three or four friends apart, BFS would find that pretty quickly. DFS, on the other hand, would go almost infinitely before it finds the match...
*/

func MaxAreaOfIsland(grid [][]int) int {
	seen := [][]bool{}
	for i := 0; i < len(grid); i++ {
		seen = append(seen, make([]bool, len(grid[i])))
	}
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			ans = int(math.Max(float64(ans), float64(Area(i, j, seen, grid))))
		}
	}
	return ans
}

func Area(r, c int, seen [][]bool, grid [][]int) int {
	if r < 0 || r >= len(seen) || c < 0 || c >= len(seen[0]) || seen[r][c] || grid[r][c] == 0 {
		return 0
	}

	seen[r][c] = true

	return 1 + Area(r+1, c, seen, grid) + Area(r-1, c, seen, grid) + Area(r, c+1, seen, grid) + Area(r, c-1, seen, grid)
}
