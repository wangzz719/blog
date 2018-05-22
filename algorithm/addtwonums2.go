package main

import "fmt"

/*
445. Add Two Numbers II : https://leetcode.com/problems/add-two-numbers-ii/description/

You are given two non-empty linked lists representing two non-negative integers. The most significant digit comes first and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Follow up:
What if you cannot modify the input lists? In other words, reversing the lists is not allowed.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNums2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	l1List := make([]int, 0)
	for l1 != nil {
		l1List = append(l1List, l1.Val)
		l1 = l1.Next
	}
	l2List := make([]int, 0)
	for l2 != nil {
		l2List = append(l2List, l2.Val)
		l2 = l2.Next
	}

	var carry int

	l1Len := len(l1List)
	l2Len := len(l2List)
	var result *ListNode

	var i, j int

	for i, j = l1Len-1, l2Len-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		fmt.Println(i, j)
		val := (l1List[i] + l2List[j] + carry) % 10
		carry = (l1List[i] + l2List[j] + carry) / 10
		if result == nil {
			result = &ListNode{
				Val:  val,
				Next: nil,
			}
		} else {
			newNode := &ListNode{
				Val:  val,
				Next: result,
			}
			result = newNode
		}
	}

	for i >= 0 {
		fmt.Println("i", i)
		val := (l1List[i] + carry) % 10
		carry = (l1List[i] + carry) / 10
		newNode := &ListNode{
			Val:  val,
			Next: result,
		}
		result = newNode
		i--
	}

	for j >= 0 {
		fmt.Println("j", j)
		val := (l2List[j] + carry) % 10
		carry = (l2List[j] + carry) / 10
		newNode := &ListNode{
			Val:  val,
			Next: result,
		}
		result = newNode
		j--
	}

	if carry > 0 {
		newNode := &ListNode{
			Val:  carry,
			Next: result,
		}
		result = newNode
	}
	return result
}
func main() {

	l1 := &ListNode{Val: 7, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}
	result := addTwoNums2(l1, l2)

	for result != nil {
		fmt.Print(result.Val, "->")
		result = result.Next
	}
	fmt.Println()
}
