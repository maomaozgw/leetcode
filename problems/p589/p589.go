package p589

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
)

type Node = treenode.NAtrTree[int]

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var result []int
	result = append(result, root.Val)
	for _, child := range root.Children {
		result = append(result, preorder(child)...)
	}
	return result
}
