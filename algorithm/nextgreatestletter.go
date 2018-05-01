package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	lettersLen := len(letters)
	l := 0
	h := lettersLen - 1
	for l <= h {
		mid := (l + h) / 2
		if letters[mid] <= target {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}
	if l < lettersLen {
		return letters[l]
	}
	return letters[0]
}

func main() {
	lts := []byte{'c', 'f', 'j'}
	fmt.Printf("%c\n", nextGreatestLetter(lts, 'd'))
}
