package main

import (
	"fmt"
	"math"
)

/*
7. Reverse Integer: https://leetcode.com/problems/reverse-integer/description/
Given a 32-bit signed integer, reverse digits of an integer.
 */
func reverse(x int) int {
	if x == 0 {
		return x
	}

	isNegative := false
	if x < 0 {
		isNegative = true
		x = -x
	}

	m := x / 10
	c := x % 10

	var result int
	for m > 0 {
		result = result*10 + c
		c = m % 10
		m = m / 10
	}

	result = result*10 + c

	if isNegative {
		result = -result
		if result < -int(math.Pow(float64(2), float64(31))) {
			return 0
		}
	}else {
		if result > int(math.Pow(float64(2), float64(31)))-1 {
			return 0
		}
	}
	
	return result
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(-120))
	fmt.Println(reverse(120))
	fmt.Println(reverse(1534236469))
}
