package main

import (
	"fmt"

	"github.com/wangzz719/blog/algorithm/model"
)

/*
2. Add Two Numbers : https://leetcode.com/problems/add-two-numbers/description/

You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/
/**
 * Definition for singly-linked list.
 * type model.ListNode struct {
 *     Val int
 *     Next *model.ListNode
 * }
 */

func addTwoNumbers(l1 *model.ListNode, l2 *model.ListNode) *model.ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l2
	}

	var result *model.ListNode
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
			result = &model.ListNode{Val: val, Next: nil}
			head = result
		} else {
			result.Next = &model.ListNode{Val: val, Next: nil}
			result = result.Next
		}
	}

	if carry != 0 {
		result.Next = &model.ListNode{Val: carry, Next: nil}
	}

	return head
}

func main() {
	l1 := &model.ListNode{Val: 2, Next: &model.ListNode{Val: 4, Next: &model.ListNode{Val: 3, Next: nil}}}
	l2 := &model.ListNode{Val: 5, Next: &model.ListNode{Val: 6, Next: &model.ListNode{Val: 4, Next: nil}}}

	result := addTwoNumbers(l1, l2)
	for result.Next != nil {
		fmt.Print(result.Val, "->")
		result = result.Next
	}

	fmt.Print(result.Val, "\n")
}
