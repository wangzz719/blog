package main

import "fmt"

// 花朵之间至少需要一个单位的间隔
// 判断是否可以种植 n 朵花
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
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
}
