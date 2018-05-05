package main

import "fmt"

func sortColors(nums []int) {
	head, tail, current := 0, len(nums)-1, 0
	for current <= tail {
		if nums[current] == 0 {
			nums[head], nums[current] = nums[current], nums[head]
			head++
			current++
		} else if nums[current] == 2 {
			nums[tail], nums[current] = nums[current], nums[tail]
			tail--
		} else {
			current++
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)

	nums = []int{2, 0, 1}
	sortColors(nums)
	fmt.Println(nums)
}
