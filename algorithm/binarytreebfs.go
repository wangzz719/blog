package main

import (
	"fmt"

	"github.com/wangzz719/blog/algorithm/model"
)

// breadth first search.
func bfs(root *model.TreeNode[any]) []any {
	if root == nil {
		return nil
	}

	nodeList := []*model.TreeNode[any]{root}
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

func toArray(root *model.TreeNode[any]) [][]any {
	if root == nil {
		return [][]any{}
	}

	result := make([][]any, 0)
	levelNodes := []*model.TreeNode[any]{root}
	for len(levelNodes) != 0 {
		levelResult := make([]any, 0)
		nextLevelNodes := make([]*model.TreeNode[any], 0)
		for _, node := range levelNodes {
			levelResult = append(levelResult, node.Val)

			if node.Left != nil {
				nextLevelNodes = append(nextLevelNodes, node.Left)
			}

			if node.Right != nil {
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
		}

		result = append(result, levelResult)
		levelNodes = nextLevelNodes
	}

	return result
}

func toArrayV2(nodes []any) [][]any {
	result := make([][]any, 0)

	if len(nodes) == 0 {
		return result
	}

	nodesCount := 1

	for len(nodes) != 0 {
		levelNodes := make([]any, 0, nodesCount)
		for j := 0; j < nodesCount; j++ {
			if nodes[j] == nil {
				continue
			}

			levelNodes = append(levelNodes, nodes[j])
		}

		nodes = nodes[nodesCount:]
		nodesCount = nodesCount * 2
		result = append(result, levelNodes)
	}

	return result
}

func main() {
	root := &model.TreeNode[any]{Val: 1}
	root.Left = &model.TreeNode[any]{Val: 2, Left: &model.TreeNode[any]{Val: 4}, Right: &model.TreeNode[any]{Val: 5}}
	root.Right = &model.TreeNode[any]{Val: 3, Left: &model.TreeNode[any]{Val: 6}, Right: &model.TreeNode[any]{Val: 7}}

	fmt.Println(bfs(root))

	fmt.Println(toArray(root))

	fmt.Println(toArrayV2([]any{1, 2, 3, nil, 5, 6, 7}))
}
