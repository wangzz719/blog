package main

import (
	"fmt"

	"github.com/wangzz719/blog/algorithm/model"
)

// breadth first search.
func bfs[T any](root *model.TreeNode[T]) []any {
	if root == nil {
		return nil
	}

	nodeList := []*model.TreeNode[T]{root}
	result := make([]any, 0)

	for len(nodeList) != 0 {
		head := nodeList[0]
		result = append(result, head.Val)

		nodeList = nodeList[1:]

		if head.Left != nil {
			nodeList = append(nodeList, head.Left)
		}

		if head.Right != nil {
			nodeList = append(nodeList, head.Right)
		}
	}

	return result
}

func main() {
	root := &model.TreeNode[int]{Val: 1}
	root.Left = &model.TreeNode[int]{Val: 2, Left: &model.TreeNode[int]{Val: 4}, Right: &model.TreeNode[int]{Val: 5}}
	root.Right = &model.TreeNode[int]{Val: 3, Left: &model.TreeNode[int]{Val: 6}, Right: &model.TreeNode[int]{Val: 7}}

	result := bfs(root)

	fmt.Println(result)
}
