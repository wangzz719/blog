package main

import (
	"fmt"

	"github.com/wangzz719/blog/model"
)

func compare(left *model.TreeNode, right *model.TreeNode) bool {
	switch {
	case left == nil && right == nil:
		return true
	case left == nil && right != nil:
		return false
	case left != nil && right == nil:
		return false
	case left.Val != right.Val:
		return false
	}

	return compare(left.Left, right.Right) && compare(left.Right, right.Left)
}

func isSymmetric(root *model.TreeNode) bool {
	if root == nil {
		return true
	}

	return compare(root.Left, root.Right)
}

func main() {
	root := &model.TreeNode{
		Val: 1,
		Left: &model.TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &model.TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println(isSymmetric(root))
}
