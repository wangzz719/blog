package main

import "fmt"

type seqLink struct {
	value int
	next  *seqLink
}

// 合并有序链表
func mergeSeqLink(link1 *seqLink, link2 *seqLink) *seqLink {
	if link1 == nil {
		return link2
	}
	if link2 == nil {
		return link2
	}

	result := &seqLink{}
	start := result
	for link1 != nil || link2 != nil {
		if link1 != nil && link2 != nil {
			if link1.value < link2.value {
				result.next = link1
				link1 = link1.next
			} else {
				result.next = link2
				link2 = link2.next
			}
		} else if link1 != nil {
			result.next = link1
			link1 = link1.next
		} else {
			result.next = link2
			link2 = link2.next
		}
		result = result.next
	}
	return start.next
}

func main() {
	link1 := &seqLink{1, &seqLink{2, &seqLink{3, nil}}}
	link2 := &seqLink{2, &seqLink{4, nil}}
	result := mergeSeqLink(link1, link2)
	for result != nil {
		fmt.Print(result.value, "==>")
		result = result.next
	}
}
