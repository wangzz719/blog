package main

import "fmt"

// 不限次买卖股票
// 在上升区间前购买，在上升区间最后卖出
func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func main() {
	fmt.Println(maxProfit([]int{1,2,4,2,5,7,2,4,9,0}))
}
