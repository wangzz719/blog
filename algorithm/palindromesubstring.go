package main

import "fmt"

func countPalindromeSubStringDP(s string) int {
	if s == "" {
		return 0
	}

	sLen := len(s)
	dp := make([][]bool, sLen)

	for i := range dp {
		dp[i] = make([]bool, sLen)
		dp[i][i] = true
	}

	count := 0

	for i := 0; i < sLen; i++ {
		for j := i; j >= 0; j-- {
			if s[i] != s[j] {
				continue
			}

			if i-j+1 <= 3 || dp[j+1][i-1] {
				dp[i][j] = true
				count++
			}
		}
	}

	return count
}

func countPalindromeSubString(s string) int {
	return len(palindromeSubString(s))
}

func palindromeSubString(s string) []string {
	if s == "" {
		return nil
	}

	sLen := len(s)
	result := make([]string, 0)

	for i := 0; i < sLen; i++ {
		for j := 0; i-j >= 0 && i+j < sLen; j++ {
			if s[i-j] != s[i+j] {
				break
			}

			result = append(result, s[i-j:i+j+1])
		}

		if i+1 < sLen && s[i] == s[i+1] {
			for j := 0; i-j >= 0 && i+1+j < sLen; j++ {
				if s[i-j] != s[i+j+1] {
					break
				}
				
				result = append(result, s[i-j:i+j+2])
			}
		}
	}

	return result
}

func main() {
	fmt.Println(palindromeSubString("abc"))
	fmt.Println(palindromeSubString("aaa"))
	fmt.Println(palindromeSubString("leetcode"))

	fmt.Println(countPalindromeSubString("leetcode"))
	fmt.Println(countPalindromeSubStringDP("leetcode"))
}
