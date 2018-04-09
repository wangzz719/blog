package main

import (
	"fmt"
	"sort"
)

// 动态规划求最长上升子数列的长度
// func longestIncrSubArrayLen(nums []int) []int 获取以 nums[i] 元素
// 结尾的最长上升子数列的长度
// length[i] = max{length[k] + 1| num[k] < num[i] (0< k < i)}
func longestIncrSubArrayLen(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	length := make([]int, len(nums), len(nums))
	length[0] = 1
	for i := 1; i < len(nums); i++ {
		maxLen := 1
		for j := 0; j < i; j++ {
			tmpLen := 1
			if nums[j] < nums[i] {
				tmpLen = length[j] + 1
			}
			if tmpLen > maxLen {
				maxLen = tmpLen
			}
		}
		length[i] = maxLen
	}
	sort.Ints(length)
	return length[len(length)-1]
}

func main() {
	length := longestIncrSubArrayLen([]int{10,9,2,5,3,7,101,18})
	fmt.Println(length)
}
