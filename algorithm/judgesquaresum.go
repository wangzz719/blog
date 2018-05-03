package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	start, end := 0, int(math.Sqrt(float64(c)))
	for start <= end {
		sum := start*start + end*end
		if sum == c {
			return true
		} else if sum > c {
			end -= 1
		} else {
			start += 1
		}
	}
	return false
}

func main() {
	fmt.Println(judgeSquareSum(5))
}
