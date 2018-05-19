package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	word1Len := len(word1)
	word2Len := len(word2)

	if word1Len == 0 || word2Len == 0 {
		return word1Len + word2Len
	}

	dp := make([][]int, word1Len+1)
	for i := 0; i <= word1Len; i++ {
		dp[i] = make([]int, word2Len+1)
		dp[i][0] = i
	}
	for i := 1; i <= word2Len; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= word1Len; i++ {
		for j := 1; j <= word2Len; j++ {
			cost := 0
			if word1[i-1] != word2[j-1] {
				cost = 1
			}
			dp[i][j] = getMinNum(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+cost)
		}
	}
	return dp[word1Len][word2Len]
}

func getMinNum(nums ...int) int {
	min := nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		}
	}
	return min
}

func main() {
	distance := minDistance("abcdef", "ef")
	fmt.Println(distance)
}
