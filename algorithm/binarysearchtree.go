package main

import (
	"fmt"

	"github.com/wangzz719/blog/algorithm/model"
)

func searchNode(root *model.TreeNode[int], val int) *model.TreeNode[int] {
	if root == nil {
		return nil
	}

	for root != nil {
		switch {
		case root.Val == val:
			return root
		case root.Val < val:
			root = root.Right
		case root.Val > val:
			root = root.Left
		}
	}

	return root
}

func insertNode(root *model.TreeNode[int], val int) {
	if root == nil {
		root = model.NewTreeNode(val)

		return
	}

	var pre *model.TreeNode[int]

	for root != nil {
		pre = root
		switch {
		case root.Val == val:
			return
		case root.Val < val:
			root = root.Right
		case root.Val > val:
			root = root.Left
		}
	}

	if pre.Val > val {
		pre.Left = model.NewTreeNode(val)

		return
	}

	pre.Right = model.NewTreeNode(val)
}

func removeNode(root *model.TreeNode[int], val int) {
	if root == nil {
		return
	}

	cur := root

	var pre *model.TreeNode[int]

	for cur != nil {
		if cur.Val == val {
			break
		}

		pre = cur

		if cur.Val < val {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}

	if cur == nil {
		return
	}

	if cur.Left == nil || cur.Right == nil {
		var child *model.TreeNode[int]

		if cur.Left != nil {
			child = cur.Left
		} else {
			child = cur.Right
		}

		if cur != root {
			if pre.Left == cur {
				pre.Left = child
			} else {
				pre.Right = child
			}
		} else {
			// 若删除节点为根节点，则重新指定根节点
			root = child
		}
	} else {
		// 获取中序遍历中待删除节点 cur 的下一个节点
		tmp := cur.Right
		for tmp.Left != nil {
			tmp = tmp.Left
		}
		// 递归删除节点 tmp
		removeNode(root, tmp.Val)
		// 用 tmp 覆盖 cur
		cur.Val = tmp.Val
	}
}

func main() {
	root := &model.TreeNode[int]{Val: 4}
	root.Left = &model.TreeNode[int]{Val: 2, Left: &model.TreeNode[int]{Val: 1}, Right: &model.TreeNode[int]{Val: 3}}
	root.Right = &model.TreeNode[int]{Val: 6, Left: &model.TreeNode[int]{Val: 5}, Right: &model.TreeNode[int]{Val: 7}}

	fmt.Println(searchNode(root, 8))
	fmt.Println(searchNode(root, 2))

	insertNode(root, 8)
	fmt.Println(searchNode(root, 8))

	removeNode(root, 9)
	fmt.Println(searchNode(root, 7))

	removeNode(root, 2)
	fmt.Println(searchNode(root, 2))
}
