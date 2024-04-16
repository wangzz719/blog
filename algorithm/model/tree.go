package model

/**
 * Definition for a binary tree node.
 */
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

type TreeNode[T any] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func newTreeNode(v any) *TreeNode[any] {
	return &TreeNode[any]{
		Left:  nil, // 左子节点指针
		Right: nil, // 右子节点指针
		Val:   v,   // 节点值
	}
}
