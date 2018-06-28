package main

import (
	"fmt"
)

func findMedianSortedArray(nums []int) float64 {
	numLen := len(nums)
	if numLen%2 == 1 {
		return float64(nums[numLen/2])
	} else {
		return float64(nums[numLen/2]+nums[numLen/2-1]) / 2
	}
}

func findKthSortedArrays(nums1 []int, l1, r1 int, nums2 []int, l2, r2, k int) int {
	if l1 > r1 {
		return nums2[l2+k-1]
	}
	if l2 > r2 {
		return nums1[l1+k-1]
	}

	mid1 := (l1 + r1) / 2
	mid2 := (l2 + r2) / 2
	if nums1[mid1] <= nums2[mid2] {
		if k <= (mid1-l1)+(mid2-l2)+1 {
			return findKthSortedArrays(nums1, l1, r1, nums2, l2, mid2-1, k)
		} else {
			return findKthSortedArrays(nums1, mid1+1, r1, nums2, l2, r2, k-(mid1-l1)-1)
		}
	} else {
		if k <= (mid1-l1)+(mid2-l2)+1 {
			return findKthSortedArrays(nums1, l1, mid1-1, nums2, l2, r2, k)
		} else {
			return findKthSortedArrays(nums1, l1, r1, nums2, mid2+1, r2, k-(mid2-l2)-1)
		}
	}

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 == 0 && len2 == 0 {
		return 0
	}
	if len1 == 0 {
		return findMedianSortedArray(nums2)
	}
	if len2 == 0 {
		return findMedianSortedArray(nums1)
	}
	/*
			int k = (m + n + 1) / 2;
		        double v = (double)FindKth(A, 0, m - 1, B, 0, n - 1, k);

		        if ((m+n) % 2 == 0) {
		            int k2 = k+1;
		            double v2 = (double)FindKth(A, 0, m - 1, B, 0, n - 1, k2);
		            v = (v + v2) / 2;
		        }

		        return v;
	*/
	k := (len1 + len2 + 1) / 2
	v := float64(findKthSortedArrays(nums1, 0, len1-1, nums2, 0, len2-1, k))
	if (len1+len2)%2 == 0 {
		v2 := float64(findKthSortedArrays(nums1, 0, len1-1, nums2, 0, len2-1, k+1))
		v = (v + v2) / 2
	}
	return v
}

func main() {
	fmt.Println(findKthSortedArrays([]int{1, 3, 4}, 0, 2, []int{2, 5, 6}, 0, 2, 3))
	fmt.Println(findMedianSortedArrays([]int{1, 3, 4}, []int{2, 5, 6}))
}
