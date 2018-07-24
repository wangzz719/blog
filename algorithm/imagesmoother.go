package main

import "fmt"

/*
661. Image Smoother: https://leetcode.com/problems/image-smoother/description/

Given a 2D integer matrix M representing the gray scale of an image,
you need to design a smoother to make the gray scale of each cell becomes the average gray scale (rounding down) of all the 8 surrounding cells and itself.
If a cell has less than 8 surrounding cells, then use as many as you can.
 */

func imageSmoother(M [][]int) [][]int {
	rLen := len(M)
	cLen := len(M[0])

	result := make([][]int, rLen)

	for i := 0; i < rLen; i++ {
		result[i] = make([]int, cLen)
	}

	for r := 0; r < rLen; r++ {
		for c := 0; c < cLen; c++ {
			count := 0
			for nr := r - 1; nr <= r+1; nr++ {
				for nc := c - 1; nc <= c+1; nc++ {
					if 0 <= nr && nr < rLen && 0 <= nc && nc < cLen {
						result[r][c] += M[nr][nc]
						count++
					}
				}
			}
			result[r][c] /= count
		}

	}

	return result
}

func main() {
	M := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	fmt.Println(imageSmoother(M))
}
