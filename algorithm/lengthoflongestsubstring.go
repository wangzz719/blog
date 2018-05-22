package main

import "fmt"

/*

3. Longest Substring Without Repeating Characters : https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	charMap := make(map[int32]int)
	length := 0
	j := 0
	for i, c := range s {
		if charMap[c] != 0 {
			if j < charMap[c] {
				j = charMap[c]
			}
		}
		charMap[c] = i + 1
		if length < (i - j + 1) {
			length = i - j + 1
		}
	}
	return length
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring("abba"))
}
