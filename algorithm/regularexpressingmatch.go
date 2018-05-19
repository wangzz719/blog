package main

import "fmt"

func isMatchRecursive(s string, p string) bool {
	if p == "" {
		if s == "" {
			return true
		} else {
			return false
		}
	}

	var firstMatch bool
	isEmpty := s == ""

	if !isEmpty && (p[0] == '.' || p[0] == s[0]) {
		firstMatch = true
	}

	if len(p) >= 2 && p[1] == '*' {
		return isMatchRecursive(s, p[2:]) || (firstMatch && isMatchRecursive(s[1:], p))
	} else {
		return firstMatch && isMatchRecursive(s[1:], p[1:])
	}
}

func main() {
	fmt.Println(isMatchRecursive("aa", "a"))
	fmt.Println(isMatchRecursive("aa", "a*"))
	fmt.Println(isMatchRecursive("aa", "a"))
	fmt.Println(isMatchRecursive("ab", ".*"))
	fmt.Println(isMatchRecursive("aab", "c*a*b"))
	fmt.Println(isMatchRecursive("aaa", "a*a*"))
}
