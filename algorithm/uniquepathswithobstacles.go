package main

import (
	"fmt"
	"strconv"
)

/*
LeetCode 63: https://leetcode.com/problems/unique-paths-ii/description/
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var pathsNum = make(map[string]int)
	if obstacleGrid[0][0] == 1 {
		pathsNum["00"] = 0
		return 0
	} else {
		pathsNum["00"] = 1
	}

	if len(obstacleGrid) <= 1 && len(obstacleGrid[0]) <= 1 {
		return pathsNum["00"]
	}

	len1 := len(obstacleGrid)
	len2 := len(obstacleGrid[0])

	for i := 1; i < len1; i++ {
		key := strconv.Itoa(i) + strconv.Itoa(0)
		keyUp := strconv.Itoa(i-1) + strconv.Itoa(0)
		if obstacleGrid[i][0] == 0 {
			pathsNum[key] = pathsNum[keyUp]
		} else {
			pathsNum[key] = 0
		}

	}
	for i := 1; i < len2; i++ {
		key := strconv.Itoa(0) + strconv.Itoa(i)
		keyRight := strconv.Itoa(0) + strconv.Itoa(i-1)
		if obstacleGrid[0][i] == 0 {
			pathsNum[key] = pathsNum[keyRight]
		} else {
			pathsNum[key] = 0
		}
	}

	var i int
	var j int
	for i = 1; i < len1; i++ {
		for j = 1; j < len2; j++ {
			key := strconv.Itoa(i) + strconv.Itoa(j)
			keyRight := strconv.Itoa(i) + strconv.Itoa(j-1)
			keyUp := strconv.Itoa(i-1) + strconv.Itoa(j)
			if obstacleGrid[i][j] == 1 {
				pathsNum[key] = 0
			} else {
				if obstacleGrid[i][j-1] == 1 && obstacleGrid[i-1][j] == 1 {
					pathsNum[key] = 0
				} else if obstacleGrid[i][j-1] == 1 {
					pathsNum[key] = pathsNum[keyUp]
				} else if obstacleGrid[i-1][j] == 1 {
					pathsNum[key] = pathsNum[keyRight]
				} else {
					pathsNum[key] = pathsNum[keyRight] + pathsNum[keyUp]
				}
			}

		}
	}
	fmt.Println(pathsNum)
	finalKey := strconv.Itoa(len1-1) + strconv.Itoa(len2-1)
	return pathsNum[finalKey]
}

func main() {
	pathNum := uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}})
	fmt.Println(pathNum)

	pathNum = uniquePathsWithObstacles([][]int{{0}})
	fmt.Println(pathNum)

	pathNum = uniquePathsWithObstacles([][]int{{0, 0}})
	fmt.Println(pathNum)

	pathNum = uniquePathsWithObstacles([][]int{{0, 1}})
	fmt.Println(pathNum)

	pathNum = uniquePathsWithObstacles([][]int{{0, 0}, {1, 1}, {0, 0}})
	fmt.Println(pathNum)

	pathNum = uniquePathsWithObstacles([][]int{{0, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}})
	fmt.Println(pathNum)
}
