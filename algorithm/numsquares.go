package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i <= n; i++ {
		for j := 1; i+j*j <= n; j++ {
			if dp[i]+1 < dp[i+j*j] {
				dp[i+j*j] = dp[i] + 1
			}
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
	fmt.Println(numSquares(3))
	fmt.Println(numSquares(18))

}
