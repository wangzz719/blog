package main

import "fmt"

// 动态规划求文本编辑距离：即让字符串 s 等于 t 需要执行的最小操作次数（替换，删除，添加）
// dp[i][j] 标识让 s[0:i+1] 等于 t[0:j+1] 时需要执行的最小操作次数
// 当 i = 0 时，d[0][j] 等于 t[0:j+1] 的长度
// 当 j = 0 时，d[i][0] 等于 d[0:i+1] 的长度
// 当 s[0:i+1] == t[0:j] 时，dp[i][j] 等于 d[i][j-1] 加上添加操作，即 dp[i][j-1] + 1
// 当 s[0:i] == t[0:j+1] 时，dp[i][j] 等于 d[i-1][j] 加上删除操作，即 dp[i-1][j] + 1
// 当 s[0:i] == t[0:j] 时，若 s[i] == t[j] 则不需要执行操作，cost = 0 ，否则需要执行替换操作，cost = 1，即 dp[i-1][j-1] + 1
// dp[i][j] = min(dp[i][j-1] + 1, dp[i-1][j] + 1, dp[i-1][j-1] + 1)
func editDistance(s string, t string) int {
	sLen := len(s)
	tLen := len(t)
	dp := make([][]int, sLen+1)
	for i := 0; i <= sLen; i++ {
		dp[i] = make([]int, tLen+1)
	}
	for i := 1; i <= sLen; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}
	for i := 1; i <= tLen; i++ {
		dp[0][i] = dp[0][i-1] + 1
	}
	for i := 1; i <= sLen; i++ {
		for j := 1; j <= tLen; j++ {
			cost := 0
			if s[i-1] != t[j-1] {
				cost = 1
			}
			dp[i][j] = getMin(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+cost)
		}
	}
	return dp[sLen][tLen]
}

func getMin(nums ...int) int {
	min := nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		}
	}
	return min
}

func main() {
	distance := editDistance("abcdef", "ef")
	fmt.Println(distance)
}
