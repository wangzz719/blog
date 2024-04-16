package main

import (
	"container/list"
	"fmt"
)

/* 基于链表实现的栈 */
type linkedListStack struct {
	// 使用内置包 list 来实现栈
	data *list.List
}

/* 初始化栈 */
func newLinkedListStack() *linkedListStack {
	return &linkedListStack{
		data: list.New(),
	}
}

/* 入栈 */
func (s *linkedListStack) push(value int) {
	s.data.PushBack(value)
}

/* 出栈 */
func (s *linkedListStack) pop() any {
	if s.isEmpty() {
		return nil
	}

	e := s.data.Back()
	s.data.Remove(e)

	return e.Value
}

/* 访问栈顶元素 */
func (s *linkedListStack) peek() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Back()
	return e.Value
}

/* 获取栈的长度 */
func (s *linkedListStack) size() int {
	return s.data.Len()
}

/* 判断栈是否为空 */
func (s *linkedListStack) isEmpty() bool {
	return s.data.Len() == 0
}

/* 获取 List 用于打印 */
func (s *linkedListStack) toList() *list.List {
	return s.data
}

func main() {
	stack := newLinkedListStack()
	fmt.Println(stack.peek())

	stack.push(1)
	stack.push(2)
	stack.push(3)

	fmt.Println(stack.peek())
	fmt.Println(stack.pop())
	fmt.Println(stack.peek())
}
