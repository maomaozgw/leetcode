package p1609

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]

func isEvenOddTree(root *TreeNode) bool {
	var nodes = []*TreeNode{root}
	var levelOdd = true
	for len(nodes) > 0 {
		var nextNodes []*TreeNode
		var prevNode = nodes[0]
		if (prevNode.Val%2 == 0) == levelOdd {
			return false
		}
		if prevNode.Left != nil {
			nextNodes = append(nextNodes, prevNode.Left)
		}
		if prevNode.Right != nil {
			nextNodes = append(nextNodes, prevNode.Right)
		}
		for i := 1; i < len(nodes); i++ {
			if (nodes[i].Val%2 == 0) == levelOdd {
				return false
			}
			diff := nodes[i].Val - prevNode.Val
			if diff == 0 {
				return false
			}
			if (diff < 0) == levelOdd {
				return false
			}
			if nodes[i].Left != nil {
				nextNodes = append(nextNodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nextNodes = append(nextNodes, nodes[i].Right)
			}
			prevNode = nodes[i]
		}
		nodes = nextNodes
		levelOdd = !levelOdd
	}
	return true
}
