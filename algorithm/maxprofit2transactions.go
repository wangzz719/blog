package main

import (
	"fmt"
	"math"
)

// 算法实现买卖两次股票，使得收益最大
// 计算当前作为第一次买时的收益
// 计算当前作为第一次卖时的收益
// 计算当前作为第二次买时的收益
// 计算当前作为第二次卖时的收益
// 结果为第二次卖时的收益
func maxProfit2transactions(prices []int) int {
	buy1 := math.MinInt64
	buy2 := math.MinInt64
	sell1 := 0
	sell2 := 0
	for _, p := range prices {
		if -p > buy1 {
			buy1 = -p
		}
		if p+buy1 > sell1 {
			sell1 = p + buy1
		}
		if sell1-p > buy2 {
			buy2 = sell1 -p
		}
		if p + buy2 > sell2 {
			sell2 = p+buy2
		}
	}
	return sell2
}

func main() {
	fmt.Println(maxProfit2transactions([]int{1,2,4,2,5,7,2,4,9,0}))
}
