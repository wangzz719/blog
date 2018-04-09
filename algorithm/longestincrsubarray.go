package main

import (
	"fmt"
	"sort"
)

// 动态规划求最长上升子数列的长度
// func longestIncrSubArrayLen(nums []int) []int 获取以 nums[i] 元素
// 结尾的最长上升子数列的长度
// length[i] = max{length[k] + 1| num[k] < num[i] (0< k < i)}
func longestIncrSubArrayLen(nums []int) []int {
	length := make([]int, len(nums), len(nums))
	length[0] = 1
	for i := 1; i < len(nums); i++ {
		maxLen := 0
		for j := 0; j < i; j++ {
			tmpLen := 0
			if nums[j] < nums[i] {
				tmpLen = length[j] + 1
			}
			if tmpLen > maxLen {
				maxLen = tmpLen
			}
		}
		length[i] = maxLen
	}
	return length
}

func main() {
	length := longestIncrSubArrayLen([]int{1, 7, 2, 8, 3, 4})
	sort.Ints(length)
	fmt.Println(length[len(length)-1])
}
