package main

import "fmt"

func searchRange(nums []int, target int) []int {
	numsLen := len(nums)
	first := binarySearch(nums, target)
	last := binarySearch(nums, target+1) - 1
	if first == numsLen || nums[first] != target {
		return []int{-1, -1}
	}
	if first > last {
		last = first
	}
	return []int{first, last}
}

func binarySearch(nums []int, target int) int {
	numsLen := len(nums)
	l, h := 0, numsLen
	for l < h {
		m := (l + h) / 2
		if nums[m] >= target {
			h = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	result := searchRange([]int{5, 7, 7, 8, 8, 8}, 8)
	fmt.Println(result)
}
