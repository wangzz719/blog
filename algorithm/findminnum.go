package main

import "fmt"

func findMin(nums []int) int {
	numsLen := len(nums)
	l, h := 0, numsLen-1
	for l < h {
		m := (l + h) / 2
		if nums[m] <= nums[h] {
			h = m
		} else {
			l = m + 1
		}
	}
	return nums[l]
}

func main() {
	num := findMin([]int{1, 2, 4, 5, 6, 7, 0})
	fmt.Println(num)
}
