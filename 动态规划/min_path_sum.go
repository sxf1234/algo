package main

import (
	"fmt"
	"math"
)

/* 最小路径和：动态规划 */
func minPathSumDP(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	dp[0][0] = grid[0][0]
	for x := 1; x < n; x++ {
		dp[x][0] = dp[x-1][0] + grid[x][0]
	}
	for y := 1; y < m; y++ {
		dp[0][y] = dp[0][y-1] + grid[0][y]
	}
	for x := 1; x < n; x++ {
		for y := 1; y < m; y++ {
			dp[x][y] = int(math.Min(float64(dp[x-1][y]), float64(dp[x][y-1]))) + grid[x][y]
		}
	}
	return dp[n-1][m-1]
}

func maxProfit(prices []int) int {
	cost := math.MaxInt
	profit := 0
	for i := 0; i < len(prices); i++ {
		cost = int(math.Min(float64(cost), float64(prices[i])))
		profit = int(math.Max(float64(profit), float64(prices[i]-cost)))
	}
	return profit
}

func main() {
	grid := [][]int{
		{1, 3, 1, 5},
		{2, 2, 4, 2},
		{5, 3, 2, 1},
		{4, 3, 5, 2},
	}

	// 动态规划
	res := minPathSumDP(grid)
	fmt.Printf("从左上角到右下角的做小路径和为  %d\n", res)
}
