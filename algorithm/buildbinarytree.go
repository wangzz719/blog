package main

import (
	"fmt"

	"github.com/wangzz719/blog/algorithm/model"
)

func buildTree(leafs []any, rootIndex int) *model.TreeNode[any] {
	if rootIndex >= len(leafs) {
		return nil
	}

	if leafs[rootIndex] == nil {
		return nil
	}

	root := model.NewTreeNode(leafs[rootIndex])
	root.Left = buildTree(leafs, 2*rootIndex+1)
	root.Right = buildTree(leafs, 2*rootIndex+2)

	return root
}

func main() {
	leafs := []any{1, 2, 3, nil, 4, 5}
	root := buildTree(leafs, 0)

	fmt.Println(root.Left)
	fmt.Println(root.Right)
}
