package main

import "fmt"

func mergeArray(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	k := m + n - 1
	for i >= 0 || j >= 0 {
		if i < 0 {
			nums1[k] = nums2[j]
			j--
		} else if j < 0 {
			nums1[k] = nums1[i]
			i--
		} else if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	mergeArray(nums1, 3, nums2, 3)
	fmt.Println(nums1)
	nums1 = []int{0}
	nums2 = []int{1}
	mergeArray(nums1, 0, nums2, 1)
	fmt.Println(nums1)
}
