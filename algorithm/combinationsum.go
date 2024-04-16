package main

import (
	"fmt"
	"sort"
)

// Given a collection of candidate numbers (candidates) and a target number (target),
// find all unique combinations in candidates where the candidate numbers sum to target.
// Each number in candidates may only be used once in the combination.
// Note: The solution set must not contain duplicate combinations.
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	path := make([]int, 0)
	length := len(candidates)
	return find(length, 0, target, 0, candidates, path)
}

func find(length, idx, target, sum int, candidates, path []int) (result [][]int) {
	if sum == target {
		found := append([]int{}, path...)
		result = append(result, found)
		return
	}

	// if idx == length {
	// 	return
	// }
	//
	// newSum := sum + candidates[idx]
	// if newSum <= target {
	// 	path = append(path, candidates[idx])
	// 	founds := find(length, idx, target, newSum, candidates, path)
	// 	result = append(result, founds...)
	// 	path = path[:len(path)-1]
	// }
	//
	// founds := find(length, idx+1, target, sum, candidates, path)
	// result = append(result, founds...)

	for i := idx; i < length; i++ {
		newSum := sum + candidates[i]
		if newSum > target {
			break
		}

		path = append(path, candidates[i])
		founds := find(length, i, target, newSum, candidates, path)
		result = append(result, founds...)

		path = path[:len(path)-1]
	}

	return
}

func main() {
	fmt.Println(combinationSum([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum([]int{8, 7, 4, 3}, 11))
}
