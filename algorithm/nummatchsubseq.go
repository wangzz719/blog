package main

import "fmt"

// 找到是给定字符串 S 子串的字符串个数
func numMatchingSubseq(S string, words []string) int {
	cnt := 0
	for _, word := range words {
		if isSubSeq(S, word) {
			cnt++
		}
	}
	return cnt
}

func isSubSeq(S string, word string) bool {
	var i, j int
	for i < len(S) && j < len(word) {
		if S[i] == word[j] {
			j++
		}
		i++
	}
	if j == len(word) {
		return true
	}
	return false
}

func main() {
	fmt.Println(numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"}))
}
