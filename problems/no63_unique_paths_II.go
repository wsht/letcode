package main

import (
	"fmt"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row, column := len(obstacleGrid), len(obstacleGrid[0])

	dp := make([][]int, row+1)
	for i := 0; i < row+1; i++ {
		dp[i] = make([]int, column+1)
	}
	//start to right
	dp[1][0] = 1

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i+1][j+1] = dp[i+1][j] + dp[i][j+1]
			}
		}
	}

	return dp[row][column]
}

func main() {
	obstacleGrid := [][]int{{0}, {0}}
	result := uniquePathsWithObstacles(obstacleGrid)
	fmt.Println(result)
}
