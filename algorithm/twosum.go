package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numsLen := len(nums)
	if numsLen < 2 {
		return nil
	}

	for start, end := 0, numsLen-1; start < end; {
		sum := nums[start] + nums[end]
		switch {
		case sum == target:
			return []int{start + 1, end + 1}
		case sum < target:
			start++
		case sum > target:
			end--
		}
	}
	
	return nil
}

func main() {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)
}
