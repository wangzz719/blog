package main

import (
	"fmt"
	"sort"
)

// Given a collection of candidate numbers (candidates) and a target number (target),
// find all unique combinations in candidates where the candidate numbers sum to target.
// Each number in candidates may only be used once in the combination.
// Note: The solution set must not contain duplicate combinations.
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	path := make([]int, len(candidates))
	return find2(candidates, 0, target, path, 0)
}

func find2(candidates []int, idx, target int, path []int, pathIdx int) (result [][]int) {
	for i := idx; i < len(candidates); i++ {
		ai := candidates[i]
		fmt.Println(idx, target, path, pathIdx)

		if i > idx && ai == candidates[i-1] {
			continue
		}

		if ai == target {
			found := make([]int, 0, pathIdx+1)
			found = append(found, path[0:pathIdx]...)
			found = append(found, ai)
			result = append(result, found)
		} else if ai < target {
			path[pathIdx] = ai
			found := find2(candidates, i+1, target-ai, path, pathIdx+1)
			result = append(result, found...)
		} else {
			break
		}
	}

	return result
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
