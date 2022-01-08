package main

import (
	"fmt"

	"github.com/wangzz719/blog/model"
)

func rightSideView(root *model.TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	search(root, &res, 0)

	return res
}

func search(node *model.TreeNode, view *[]int, depth int) {
	if node == nil {
		return
	}

	if len(*view) == depth {
		*view = append(*view, node.Val)
	}

	search(node.Right, view, depth + 1)
	search(node.Left, view, depth + 1)
}

func main() {
	root := &model.TreeNode{
		Val: 1,
		Left: &model.TreeNode{
			Val: 2,
			Left: &model.TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &model.TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: nil,
	}

	fmt.Println(rightSideView(root))
}
