package main

import (
	"sort"
	"fmt"
)

// 实现排序接口
type Points [][]int

func (p Points) Len() int      { return len(p) }
func (p Points) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Points) Less(i, j int) bool {
	return p[i][1] < p[j][1]
}

// 贪心算法实现找到刺破所有气球的最小发射飞镖数
// 每次射出一个飞镖时都保证左边已经没有气球，同时此次飞镖能刺破最多的气球
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Sort(Points(points))
	fmt.Println(points)
	cp := points[0][1]
	ac := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] <= cp {
			continue
		}
		cp = points[i][1]
		ac ++
	}
	return ac
}

func main() {
	c := findMinArrowShots([][]int{{1, 6}, {7, 12}, {2, 8}, {10, 16}})
	fmt.Println(c)
}
