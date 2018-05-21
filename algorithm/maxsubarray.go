package main

import (
	"fmt"
	"math"
)

/*
53. Maximum Subarray: https://leetcode.com/problems/maximum-subarray/description/

Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
*/

func maxSubArray(nums []int) int {
	numsLen := len(nums)
	if numsLen == 1 {
		return nums[0]
	}
	dp := make([]int, numsLen, numsLen) // 最大子串是以当前元素结尾的子串和
	dp[0] = nums[0]
	for i := 1; i < numsLen; i++ {
		dp[i] = nums[i]
		if dp[i-1]+nums[i] > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		}
	}

	res := math.MinInt32
	for i := 0; i < numsLen; i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
