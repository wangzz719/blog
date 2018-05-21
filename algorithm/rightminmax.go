package main

import "fmt"

/*
单调栈：求数组中每个元素右边第一个比该元素大的数字，没有时为 -1
*/
func rightMinMax(nums []int) []int {
	stack := []int{-1}
	res := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		j := len(stack) - 1
		for ; j >= 0; j-- {
			if stack[j] == -1 || stack[j] > nums[i] {
				break
			}
		}
		res = append(res, stack[j])
		stack = append(stack[0:j+1], nums[i])
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func main() {
	fmt.Println(rightMinMax([]int{9, 4, 5, 7, 3, 2, 10}))
}
