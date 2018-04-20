package main

import "fmt"

// 快速排序
func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	start, end := 0, len(nums)-1
	mid, current := nums[0], 1
	for start < end {
		if nums[current] > mid {
			nums[current], nums[end] = nums[end], nums[current]
			end--
		} else {
			nums[start], nums[current] = nums[current], nums[start]
			start++
			current++
		}
	}
	nums[start] = mid
	quickSort(nums[:start])
	quickSort(nums[start+1:])
	return nums
}

func main() {
	fmt.Println(quickSort([]int{2, 1, 3, 4, 0}))
}
