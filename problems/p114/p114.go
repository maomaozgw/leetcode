package p114

import (
	"fmt"
	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]

func flatten(root *TreeNode) {
	var s = treeToSlice(root)

	current := root
	for _, node := range s {
		fmt.Println(node.Val)
		current.Left = nil
		current.Right = node
		current = node
	}
}

func treeToSlice(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	var s []*TreeNode
	if root.Left != nil {
		s = append(s, root.Left)
		leftNodes := treeToSlice(root.Left)
		if len(leftNodes) > 0 {
			s = append(s, leftNodes...)
		}
	}
	if root.Right != nil {
		s = append(s, root.Right)
		rightNodes := treeToSlice(root.Right)
		if len(rightNodes) > 0 {
			s = append(s, rightNodes...)
		}
	}
	return s
}

func flattenO1(root *TreeNode) {
	if root.Left != nil {
		flattenO1(root.Left)
		tmp := root.Right
		root.Right = root.Left
		root.Left = nil
		flattenO1(tmp)
		root.Right.Right = tmp

	} else if root.Right != nil {
		flattenO1(root.Right)
	}
}
