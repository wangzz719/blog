package main

import "fmt"

/*
2. Add Two Numbers : https://leetcode.com/problems/add-two-numbers/description/

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l2
	}
	var result *ListNode
	head := result

	carry := 0
	for l1 != nil || l2 != nil {
		var val int
		if l1 != nil && l2 != nil {
			val = (l1.Val + l2.Val + carry) % 10
			carry = (l1.Val + l2.Val + carry) / 10
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			val = (l1.Val + carry) % 10
			carry = (l1.Val + carry) / 10

			l1 = l1.Next
		} else if l2 != nil {
			val = (l2.Val + carry) % 10
			carry = (l2.Val + carry) / 10
			l2 = l2.Next
		}
		if result == nil {
			result = &ListNode{
				Val:  val,
				Next: nil,
			}
			head = result
		} else {
			result.Next = &ListNode{
				Val:  val,
				Next: nil,
			}
			result = result.Next
		}
	}

	if carry != 0 {
		result.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	return head
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}
	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Print(result.Val, "->")
		result = result.Next
	}
	fmt.Println()
}
