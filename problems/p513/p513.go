package p513

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func findBottomLeftValue(root *TreeNode) int {
	var result = 0 // 记录最深的左节点
	var nodes = []*TreeNode{root}
	for len(nodes) > 0 {
		result = nodes[0].Val
		var nextNodes []*TreeNode
		for _, node := range nodes {
			if node.Left != nil {
				nextNodes = append(nextNodes, node.Left)
			}
			if node.Right != nil {
				nextNodes = append(nextNodes, node.Right)
			}
		}
		nodes = nextNodes
	}
	return result
}
