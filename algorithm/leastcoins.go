package main

import "fmt"

// 动态规划求解最少硬币数满足凑够 sum
// 总和为 i 时需要的最少硬币数为 i 减去硬币面值时需要的最少硬币数加 1
// l[i] = min {l[i-k] + 1| i-k >0 and k in CoinValues}
// example:
// sum = 11, conValues = []int{1, 3, 5}
// l[0] = 0, l[1] = min{l[1-1] + 1}
// l[2] = min{l[2-1] + 1}, l[3] = min{l[3-1] + 1, l[3-3]+ 1}
// ...

func leastCoins(sum int, coinValues []int) ([]int, int) {
	sumCoins := make(map[int]int)
	result := make([]int, 0)
	count := make([]int, sum+1, sum+1)

	count[0] = 0
	for i := 1; i <= sum; i++ {
		minCount := sum
		coin := -1
		for _, v := range coinValues {
			tmp := 0
			if i-v >= 0 {
				tmp = count[i-v] + 1
				if tmp < minCount {
					minCount = tmp
					coin = v
				}
			}
		}

		count[i] = minCount
		sumCoins[i] = coin
	}

	fmt.Println(sumCoins)
	fmt.Println(count)

	for sum > 0 {
		result = append(result, sumCoins[sum])
		sum -= sumCoins[sum]
	}

	return result, len(result)
}

func main() {
	fmt.Println(leastCoins(11, []int{1, 3, 5}))
}
