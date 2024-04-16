package main

import "fmt"

// 假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
// 给你一个整数数组 flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。
// 另有一个数 n ，能否在不打破种植规则的情况下种入 n 朵花？能则返回 true ，不能则返回 false 。
func canPlaceFlowers(flowerbed []int, n int) bool {
	cnt := 0
	length := len(flowerbed)
	for i := 0; i < length; i++ {
		if flowerbed[i] == 1 {
			continue
		}

		prev := 0
		if i != 0 {
			prev = flowerbed[i-1]
		}

		next := 0
		if i != length-1 {
			next = flowerbed[i+1]
		}

		if prev == 0 && next == 0 {
			cnt++
			flowerbed[i] = 1
		}
	}

	return cnt >= n
}

func main() {
	fmt.Println(canPlaceFlowers(nil, 1))
	fmt.Println(canPlaceFlowers([]int{}, 1))
	fmt.Println(canPlaceFlowers([]int{0}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
}
