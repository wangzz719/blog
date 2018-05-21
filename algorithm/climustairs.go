package main

import "fmt"

/*
70. Climbing Stairs : https://leetcode.com/problems/climbing-stairs/description/

You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

*/
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n, n) // dp 表示上到 n 层台阶的方法数
	dp[0] = 1
	dp[1] = 2

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

func main() {
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(4))
}
