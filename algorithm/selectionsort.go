package main

import "fmt"

func selectionSort(nums []int) {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}

		nums[i], nums[min] = nums[min], nums[i]
	}
}

func main() {
	nums := []int{3, 2, 5, 8, 1}
	selectionSort(nums)
	fmt.Println(nums)
}
