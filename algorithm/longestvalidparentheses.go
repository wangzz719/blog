package main

import "fmt"

/*
动态规划求最长匹配的括号
dp[i]标识以 s[i] 结尾的匹配的括号字符串长度，只有当 dp[i] == ')' 时才会可能产生新的匹配括号组，公式如下：

if s[i-1] == '(' :
	dp[i] = dp[i-2] + 2 if i >= 2 else dp[i] = 2
else:
	if s[i-dp[i-1]-1] == '(' and i-dp[i-1]-1 >= 0:
		if i-dp[i-1] >= 2 :
			dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2 	tips: "(,sub_s,)"
		else :
			dp[i] = dp[i-1] + 2
*/
func longestValidParentheses(s string) int {
	if s == "" {
		return 0
	}
	dp := make([]int, len(s), len(s))

	var maxLen int

	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				var tmp int
				if i-dp[i-1] >= 2 {
					tmp = dp[i-dp[i-1]-2]
				}
				dp[i] = dp[i-1] + tmp + 2
			}
			if maxLen < dp[i] {
				maxLen = dp[i]
			}
		}
	}
	return maxLen
}

func isValidParentheses(s string) bool {
	parenthess := map[int32]int32{')': '(', ']': '[', '}': '{'}
	stack := make([]int32, 0)
	if len(s)%2 != 0 {
		return false
	}
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else if len(stack) > 0 && parenthess[c] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(longestValidParentheses("())"))
	fmt.Println(isValidParentheses("()"))
	fmt.Println(isValidParentheses("([)]"))
}
