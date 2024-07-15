package p2196

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodeMap := map[int]*TreeNode{}
	hasParent := map[int]bool{}
	for i := range descriptions {
		parent, child, isLeft := descriptions[i][0], descriptions[i][1], descriptions[i][2]
		if _, ok := nodeMap[parent]; !ok {
			nodeMap[parent] = &TreeNode{Val: parent}
		}
		hasParent[child] = true
		if _, ok := nodeMap[child]; !ok {
			nodeMap[child] = &TreeNode{Val: child}
		}
		if isLeft == 1 {
			nodeMap[parent].Left = nodeMap[child]
		} else {
			nodeMap[parent].Right = nodeMap[child]
		}
	}
	var root *TreeNode
	for i := range descriptions {
		if !hasParent[descriptions[i][0]] {
			return nodeMap[descriptions[i][0]]
		}
	}
	return root
}
