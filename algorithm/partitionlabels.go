package main

import "fmt"

// 贪心算法实现：分隔字符串使同种字符出现在一起单独的子串中，返回每个子串的长度
func partitionLabels(S string) []int {
	lastPosition := make(map[int32]int)
	for i, c := range S {
		lastPosition[c] = i
	}

	var j, anchor int
	result := make([]int, 0)
	for i, c := range S {
		p := lastPosition[c]
		if j < p {
			j = p
		}
		if j == i {
			result = append(result, i-anchor+1)
			anchor = i + 1
		}
	}
	return result
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}
