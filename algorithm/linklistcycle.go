package main

import "fmt"

type Node struct {
	val  int64
	next *Node
}

func hasCycle(n *Node) bool {
	if n == nil || n.next == nil {
		return false
	}

	slow := n
	fast := n

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			return true
		}
	}

	return false
}

func findCycleStartNode(n *Node) *Node {
	if n == nil || n.next == nil {
		return nil
	}

	slow := n
	fast := n

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			ptr := n

			for ptr != slow {
				ptr = ptr.next
				slow = slow.next
			}

			return ptr
		}
	}

	return nil
}

func main() {
	n1 := &Node{val: 1}
	n2 := &Node{val: 2}
	n3 := &Node{val: 3}
	n4 := &Node{val: 4}

	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n2

	fmt.Println(hasCycle(n1))
	fmt.Printf("node start: %+v\n", findCycleStartNode(n1))
}
