package main

import "fmt"

/*
62. Unique Paths : https://leetcode.com/problems/unique-paths/description/

A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?
*/

func uniquePaths(m int, n int) int {
	if m == 0 && n == 0 {
		return 0
	}

	dp := make([][]int, m, m) // dp 表示 m * n 矩阵到达终点的路径数
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(7, 3))
}
