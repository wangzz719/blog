package main

import "fmt"

func isSubsequence(s string, t string) bool {
	var i, j int
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	if i == len(s) {
		return true
	}
	return false
}

func main() {
	fmt.Println(isSubsequence("abc", "ahgdc"))
}
