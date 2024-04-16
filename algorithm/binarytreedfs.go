package main

import (
	"fmt"

	"github.com/wangzz719/blog/algorithm/model"
)

// pre-order depth first search.
func preOrderDFS[T any](root *model.TreeNode[T]) []any {
	if root == nil {
		return nil
	}

	result := []any{root.Val}
	result = append(result, preOrderDFS(root.Left)...)
	result = append(result, preOrderDFS(root.Right)...)

	return result
}

// in-order depth first search
func inOrderDFS[T any](root *model.TreeNode[T]) []any {
	if root == nil {
		return nil
	}

	result := make([]any, 0)
	result = append(result, inOrderDFS(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inOrderDFS(root.Right)...)

	return result
}

// post-order depth first search
func postOrderDFS[T any](root *model.TreeNode[T]) []any {
	if root == nil {
		return nil
	}

	result := make([]any, 0)
	result = append(result, postOrderDFS(root.Left)...)
	result = append(result, postOrderDFS(root.Right)...)
	result = append(result, root.Val)

	return result
}

func main() {
	root := &model.TreeNode[int]{Val: 1}
	root.Left = &model.TreeNode[int]{Val: 2, Left: &model.TreeNode[int]{Val: 4}, Right: &model.TreeNode[int]{Val: 5}}
	root.Right = &model.TreeNode[int]{Val: 3, Left: &model.TreeNode[int]{Val: 6}, Right: &model.TreeNode[int]{Val: 7}}

	result := preOrderDFS(root)
	fmt.Println(result)

	result = inOrderDFS(root)
	fmt.Println(result)

	result = postOrderDFS(root)
	fmt.Println(result)
}
