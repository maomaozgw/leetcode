package p590

import "github.com/maomaozgw/leetcode/structures/treenode"

type Node = treenode.NAtrTree[int]

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var result []int
	for _, child := range root.Children {
		result = append(result, postorder(child)...)
	}
	result = append(result, root.Val)
	return result
}
