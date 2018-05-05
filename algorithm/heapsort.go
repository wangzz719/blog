package main

import "fmt"

// 调整为大顶堆
func maxHeapify(nums []int, start int, end int) {
	son := 2*start + 1
	dad := start
	for son <= end {
		if son+1 <= end && nums[son] < nums[son+1] { // 先比较两个子节点，选择较大的
			son++
		}
		if nums[dad] > nums[son] { // 如果父节点大于子节点，则堆调整完毕
			return
		} else { // 交换父节点和子节点，继续调整子孙节点
			nums[dad], nums[son] = nums[son], nums[dad]
			dad = son
			son = dad*2 + 1
		}
	}
}

// 堆排序
func heapSort(nums []int) {
	numsLen := len(nums)

	// 初始化，从最后一个父节点开始调整
	for i := numsLen/2 - 1; i >= 0; i-- {
		maxHeapify(nums, i, numsLen-1)
	}
	//先取出堆顶元素和最后一个元素交换，再调整剩余元素为大顶堆，直到排序完毕
	for i := numsLen - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		maxHeapify(nums, 0, i-1)
	}
}

func main() {
	nums := []int{3, 5, 3, 0, 8, 6, 1, 5, 8, 6, 2, 4, 9, 4, 7, 0, 1, 8, 9, 7, 3, 1, 2, 5, 9, 7, 4, 0, 2, 6}
	heapSort(nums)
	fmt.Println(nums)

}
