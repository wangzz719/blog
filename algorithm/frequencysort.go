package main

import (
	"bytes"
	"fmt"
)

func frequencySort(s string) string {
	frequentChar, max := getCharCount(s)
	result := ""
	for max > 0 {
		for _, b := range frequentChar[max] {
			result += getString(b, max)
		}
		max--
	}
	return result
}

func getString(b byte, frequency int) string {
	s := bytes.Repeat([]byte{b}, frequency)
	return string(s)
}

func getCharCount(s string) (map[int][]byte, int) {
	charCount := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		charCount[s[i]] += 1
	}
	frequentChar := make(map[int][]byte)
	max := 0
	for c, v := range charCount {
		frequentChar[v] = append(frequentChar[v], c)
		if v > max {
			max = v
		}
	}
	return frequentChar, max
}

func main() {
	fmt.Println(frequencySort("Aabb"))
}
