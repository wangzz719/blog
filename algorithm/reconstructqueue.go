package main

import (
	"fmt"
	"sort"
)

type People [][]int

func (s People) Len() int {
	return len(s)
}

func (s People) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s People) Less(i, j int) bool {
	if s[i][0] == s[j][0] {
		return s[i][1] < s[j][1]
	} else {
		return s[i][0] > s[j][0]
	}
}

func reconstructQueue(people [][]int) [][]int {
	if len(people) == 0 || len(people[0]) == 0 {
		return [][]int{}
	}
	sort.Sort(People(people))
	result := make([][]int, 0)
	for _, p := range people {
		fmt.Println(p)
		tmp := append([][]int{}, result[p[1]:]...)
		result = append(result[:p[1]], p)
		result = append(result, tmp...)
	}
	return result
}

func main() {
	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
}
