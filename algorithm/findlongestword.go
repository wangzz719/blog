package main

import (
	"fmt"
	"strings"
)

func findLongestWord(s string, d []string) string {
	longestWord := ""
	for _, target := range d {
		l1 := len(longestWord)
		l2 := len(target)
		if l1 > l2 || (l1 == l2 && strings.Compare(longestWord, target) < 0) {
			continue
		}
		if isValid(s, target) {
			longestWord = target
		}
	}
	return longestWord
}

func isValid(s string, target string) bool {
	sLen := len(s)
	tLen := len(target)
	var i, j int
	for i < sLen && j < tLen {
		if s[i] == target[j] {
			j++
		}
		i++
	}
	return j == tLen
}

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
	fmt.Println(findLongestWord("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
}
