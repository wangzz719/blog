package main

import "fmt"

// 判断一个数组能不能只修改一个数就成为非递减数组
// 在出现 nums[i] < nums[i - 1] 时，需要考虑的是应该修改数组的哪个数，
// 使得本次修改能使 i 之前的数组成为非递减数组，并且 不影响后续的操作
func checkPossibility(nums []int) bool {
	cnt := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			cnt++
			if i-2 >= 0 && nums[i-2] > nums[i] {
				nums[i] = nums[i-1]
			} else {
				nums[i-1] = nums[i]
			}
		}
	}
	if cnt > 1 {
		return false
	}
	return true
}

func main() {
	fmt.Println(checkPossibility([]int{3, 4, 2, 3}))
	fmt.Println(checkPossibility([]int{4, 2, 3}))
}
