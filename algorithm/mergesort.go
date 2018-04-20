package main

import "fmt"

// 归并排序
func mergeSort(nums []int) []int {
	numsLen := len(nums)
	if numsLen <= 1 {
		return nums
	}
	mid := numsLen / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	leftLen := len(left)
	rightLen := len(right)

	var i, j int
	result := make([]int, 0)
	for i < leftLen && j < rightLen {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func main() {
	fmt.Println(mergeSort([]int{2, 1, 3, 0, 4}))
}
