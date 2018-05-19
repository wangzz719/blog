package main

import (
	"fmt"
	"strconv"
)

func minPathSum(grid [][]int) int {
	var sum = make(map[string]int)
	sum["00"] = grid[0][0]
	for i, g1 := range grid {
		for j, g2 := range g1 {
			key := strconv.Itoa(i) + strconv.Itoa(j)
			keyRight := strconv.Itoa(i) + strconv.Itoa(j-1)
			keyUp := strconv.Itoa(i-1) + strconv.Itoa(j)

			rightSum := sum[keyRight]

			upSum := sum[keyUp]

			if j-1 >= 0 && i-1 >= 0 {
				if rightSum+g2 < upSum+g2 {
					sum[key] = rightSum + g2
				} else {
					sum[key] = upSum + g2
				}
			} else if j-1 >= 0 {
				sum[key] = rightSum + g2
			} else if i-1 >= 0 {
				sum[key] = upSum + g2
			}
		}
	}
	finalKey := strconv.Itoa(len(grid)-1) + strconv.Itoa(len(grid[0])-1)
	return sum[finalKey]
}

func main() {
	sum := minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})
	fmt.Println(sum)

	sum = minPathSum([][]int{{0, 2, 2, 6, 4, 1, 6, 2, 9, 9, 5, 8, 4, 4}, {0, 3, 6, 4, 5, 5, 9, 7, 8, 3, 9, 9, 5, 4}, {6, 9, 0, 7, 2, 2, 5, 6, 3, 1, 0, 4, 2, 5}, {3, 8, 2, 3, 2, 8, 8, 7, 5, 9, 6, 3, 4, 5}, {4, 0, 1, 3, 9, 2, 0, 1, 6, 7, 9, 2, 8, 9}, {6, 2, 7, 9, 0, 9, 5, 2, 7, 5, 1, 4, 4, 7}, {9, 8, 3, 3, 0, 6, 8, 0, 8, 8, 3, 5, 7, 3}, {7, 7, 4, 5, 9, 1, 5, 0, 2, 2, 2, 1, 7, 4}, {5, 1, 3, 4, 1, 6, 0, 4, 3, 8, 4, 3, 9, 9}, {0, 6, 4, 9, 4, 1, 5, 5, 4, 2, 5, 7, 4, 0}, {0, 1, 6, 6, 4, 9, 2, 7, 8, 2, 1, 3, 3, 7}, {8, 4, 9, 9, 2, 3, 8, 6, 6, 5, 4, 1, 7, 9}})
	fmt.Println(sum)
}
