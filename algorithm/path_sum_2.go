package main

import (
	"fmt"

	"github.com/wangzz719/blog/model"
)

/**
 * Definition for a binary tree node.
 */

func pathSum(root *model.TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	tmp := make([]int, 0)
	dfs(root, targetSum, &res, tmp)

	return res
}

func dfs(root *model.TreeNode, targetSum int, res *[][]int, carry []int) {
	carry = append(carry, root.Val)
	targetSum = targetSum - root.Val

	if root.Left == nil && root.Right == nil {
		if targetSum == 0 {
			tmp := make([]int, len(carry))

			copy(tmp, carry)

			*res = append(*res, tmp)
		}

		return
	}

	if root.Left != nil {
		dfs(root.Left, targetSum, res, carry)
	}

	if root.Right != nil {
		dfs(root.Right, targetSum, res, carry)
	}
}


func main() {
	root := &model.TreeNode{
		Val: 5,
		Left: &model.TreeNode{
			Val: 4,
			Left: &model.TreeNode{
				Val: 11,
				Left: &model.TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &model.TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &model.TreeNode{
			Val: 8,
			Left: &model.TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &model.TreeNode{
				Val: 4,
				Left: &model.TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: &model.TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	fmt.Println(pathSum(root, 22))
}
