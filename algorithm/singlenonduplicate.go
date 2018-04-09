package main

import "fmt"

// 二分查找寻找没有重复的唯一数字
// 思路：如果重复元素都出现两次的话，那么唯一没有出现重复的元素会出现在偶数位（以 0 开始）
// 没有重复的元素会打破 nums[2n] == nums[2n+1] 的关系
// 采用二分查找，如果 mid 为中间的偶数位，如果 nums[mid] == nums[mid+1] 说明没有重复的元素出现在 mid 之后，
// 反之出现在 mid 之前

func singleNonDuplicate(nums []int) int {
	l, h := 0, len(nums)-1
	for l < h {
		mid := (l + h) / 2
		if mid%2 == 1 {
			mid -= 1
		}
		if nums[mid] == nums[mid+1] {
			l = mid + 2
		} else {
			h = mid
		}
	}
	return nums[l]
}

func main() {
	fmt.Println(singleNonDuplicate([]int{1,1,2,3,3,4,4,8,8}))
}
