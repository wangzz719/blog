package main

import "fmt"

func binarySearch(nums []int, target int) int {
	i, j := 0, len(nums)-1

	for i <= j {
		m := (i + j) / 2

		if nums[m] < target { // 此情况说明 target 在区间 [m+1, j] 中
			i = m + 1
		} else if nums[m] > target { // 此情况说明 target 在区间 [i, m-1] 中
			j = m - 1
		} else { // 找到目标元素，返回其索引
			return m
		}
	}

	return -1
}

func binarySearchRotateArray(nums []int, target int) int {
	i := 0

	for i < len(nums)-1 {
		if nums[i] > nums[i+1] {
			break
		}

		i++
	}

	if target >= nums[0] {
		return binarySearch(nums[0:i+1], target)
	}

	res := binarySearch(nums[i+1:], target)
	if res == -1 {
		return res
	}

	return i + res + 1
}

func main() {
	fmt.Println(binarySearch([]int{}, 3))
	fmt.Println(binarySearch([]int{1}, 1))
	fmt.Println(binarySearch([]int{1, 2, 4, 5, 9}, 9))

	fmt.Println(binarySearchRotateArray([]int{4, 5, 6, 7, 0, 1, 2}, 9))
	fmt.Println(binarySearchRotateArray([]int{1}, 1))
	fmt.Println(binarySearchRotateArray([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(binarySearchRotateArray([]int{3, 1}, 1))
}
