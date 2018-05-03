package main

import "fmt"

var vowels = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}

func reverseVowels(s string) string {
	start, end := 0, len(s)-1
	result := make([]byte, len(s))
	for start <= end {
		if !vowels[s[start]] {
			result[start] = s[start]
			start++
		} else if !vowels[s[end]] {
			result[end] = s[end]
			end--
		} else {
			result[start], result[end] = s[end], s[start]
			start++
			end--
		}
	}
	return string(result)
}

func main() {
	fmt.Println(reverseVowels("leetcode"))
	fmt.Println(reverseVowels("hello"))
	fmt.Printf("%T", reverseVowels(" "))
}
