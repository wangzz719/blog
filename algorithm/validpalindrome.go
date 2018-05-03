package main

import "fmt"

func validPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start < end {
		if s[start] != s[end] {
			return isPalindrome(s, start+1, end) || isPalindrome(s, start, end-1)
		}
		start++
		end--
	}
	return true
}

func isPalindrome(s string, start int, end int) bool {
	for start < end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			return false
		}

	}
	return true
}

func main() {
	fmt.Println(validPalindrome("abca"))
}
