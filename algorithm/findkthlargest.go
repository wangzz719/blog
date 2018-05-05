package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	nums = quickSort2(nums)
	return nums[len(nums)-k]
}

func quickSort2(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	head, tail := 0, len(nums)-1
	current := 1
	mid := nums[0]

	for head < tail {
		if nums[current] < mid {
			nums[current], nums[head] = nums[head], nums[current]
			current++
			head++
		} else {
			nums[current], nums[tail] = nums[tail], nums[current]
			tail--
		}
	}
	nums[head] = mid
	quickSort2(nums[:head])
	quickSort2(nums[head+1:])
	return nums
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(nums, 2))

	nums2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(nums2, 4))

	nums3 := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(nums3, 2))

}
