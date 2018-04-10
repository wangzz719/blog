package main

import (
	"math"
	"fmt"
)

// 算法实现只买卖一次股票，实现收益最大
// 判断当前购买时的收益和当前卖出时的收益
func maxProfit1transactions(prices []int) int {
	buy := math.MinInt64
	sell := 0
	for _, p := range prices {
		if -p > buy {
			buy = -p
		}
		if p+buy > sell {
			sell = p + buy
		}
	}
	return sell
}

// 算法实现只买卖一次股票，实现收益最大
// 两个游标，当卖点比买入点低时，以卖点作为买入点，重新计算收益
// 记录中间状态最大收益
func maxProfit1transactions2(prices []int) int {
	var begin, end, max int
	for end < len(prices) {
		profit := prices[end] - prices[begin]
		if profit <= 0 {
			begin = end
		} else if profit > max {
			max = profit
		}
		end ++
	}
	return max
}

func main() {
	fmt.Println(maxProfit1transactions([]int{9, 2, 4, 2, 5, 17, 2, 4, 9, 0}))
	fmt.Println(maxProfit1transactions2([]int{9, 2, 4, 2, 5, 17, 2, 4, 9, 0}))
}
