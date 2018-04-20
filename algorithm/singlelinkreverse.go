package main

import "fmt"

type singleLink struct {
	value int
	next  *singleLink
}

// 反转单链表
func reverseSingleLink(link *singleLink) *singleLink {
	if link == nil || link.next == nil {
		return link
	}

	var pre, pnext *singleLink
	pnow := link

	var result *singleLink

	for pnow != nil {
		pnext = pnow.next
		if pnext == nil {
			result = pnow
		}
		pnow.next = pre
		pre = pnow
		pnow = pnext
	}
	return result
}

func main() {
	link := &singleLink{1, &singleLink{2, &singleLink{3, nil}}}
	reversedLink := reverseSingleLink(link)
	for reversedLink != nil {
		fmt.Print(reversedLink.value, "==>")
		reversedLink = reversedLink.next
	}
}
