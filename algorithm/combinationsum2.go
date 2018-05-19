package main

import (
	"fmt"
	"sort"
)

func combinationSum2(a []int, target int) [][]int {
	sort.Ints(a)
	path := make([]int, len(a))
	return find(a, 0, target, path, 0)
}

func find(a []int, idx, target int, path []int, pathIdx int) (result [][]int) {
	for i := idx; i < len(a); i++ {
		ai := a[i]
		if i > idx && ai == a[i-1] {
			continue
		}
		if ai == target {
			found := make([]int, 0, pathIdx+1)
			found = append(found, path[0:pathIdx]...)
			found = append(found, ai)
			result = append(result, found)
		} else if ai < target {
			path[pathIdx] = ai
			found := find(a, i+1, target-ai, path, pathIdx+1)
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
