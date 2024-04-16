package main

import (
	"fmt"
	"math"

	"github.com/wangzz719/blog/algorithm/model"
)

func buildTree(leafs []int, rootIndex int) *model.TreeNode[any] {
	if rootIndex >= len(leafs) {
		return nil
	}

	if leafs[rootIndex] == math.MinInt {
		return nil
	}

	root := model.NewTreeNode(leafs[rootIndex])
	root.Left = buildTree(leafs, 2*rootIndex+1)
	root.Right = buildTree(leafs, 2*rootIndex+2)

	return root
}

func main() {
	leafs := []int{1, 2, 3, math.MinInt, 4, 5}
	root := buildTree(leafs, 0)

	fmt.Println(root.Left)
	fmt.Println(root.Right)
}
