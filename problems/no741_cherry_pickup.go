package main

import (
	"fmt"
	"math"
)

//这里对于断路的处理方法，以及最长路径回溯的处理方式有一些问题。
func cherryPickup(grid [][]int) int {
	row := len(grid)
	column := len(grid[0])
	total := 0

	dp := make([][]int, row+1)
	for run := 0; run < 2; run++ {
		//初始化dp
		for index := range dp {
			dp[index] = make([]int, column+1)
		}
		for i := 0; i < row; i++ {
			for j := 0; j < column; j++ {
				if (grid[i][j]) != -1 {
					top := dp[i][j+1]
					left := dp[i+1][j]

					//check grid top and left both -1
					if (i > 0 && j > 0 && grid[i-1][j] == -1 && grid[i][j-1] == -1) || (i > 0 && j == 0 && grid[i-1][j] == -1) || (i == 0 && j > 0 && grid[i][j-1] == -1) {
						fmt.Println(i, j)
						grid[i][j] = -1
						dp[i+1][j+1] = 0
					} else {
						dp[i+1][j+1] = int(math.Max(float64(top), float64(left))) + grid[i][j]
					}
				} else {
					dp[i+1][j+1] = 0
				}
			}
		}

		total += dp[row][column]
		//如果第一次最大为0那么，没有继续的必要了
		if total == 0 && run == 0 {
			return total
		}
		//update grid with long path
		grid[row-1][column-1] = 0
		grid[0][0] = 0
		for i := row; i > 1; i-- {
			for j := column; j > 1; j-- {
				if dp[i][j-1] >= dp[i-1][j] {
					grid[i-1][j-2] = 0
				} else {
					grid[i-2][j-1] = 0
				}
			}
		}
	}
	return total
}

func main() {
	grid := [][]int{{1, 1, -1}, {1, -1, 1}, {-1, 1, 1}}
	fmt.Println(cherryPickup(grid))
}
