package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	numCount := getNumCount(nums)
	numFrequent := make(map[int][]int)
	max := 0
	for k, v := range numCount {
		if v > max {
			max = v
		}
		numFrequent[v] = append(numFrequent[v], k)
	}

	result := make([]int, 0)
	for len(result) < k {
		resultLen := len(result)
		if resultLen+len(numFrequent[max]) < k {
			result = append(result, numFrequent[max]...)
		} else {
			result = append(result, numFrequent[max][:k-resultLen]...)
		}
		max--
	}
	return result
}

func getNumCount(nums []int) map[int]int {
	result := make(map[int]int)
	for _, num := range nums {
		result[num] += 1
	}
	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(topKFrequent(nums, 2))
}
