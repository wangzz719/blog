package main

import "fmt"

func bubbleSort(nums []int) {
	length := len(nums)

	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func main() {
	nums := []int{3, 2, 5, 8, 1}
	bubbleSort(nums)
	fmt.Println(nums)
}
