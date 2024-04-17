package p988

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]

func smallestFromLeaf(root *TreeNode) string {
	var result = "{"
	var dfs func(node *TreeNode, path []byte)

	dfs = func(node *TreeNode, path []byte) {
		if node == nil {
			return
		}
		path = append(path, byte(node.Val+'a'))
		if node.Left == nil && node.Right == nil {
			pathResult := string(reverse(path))

			if pathResult < result {
				result = pathResult
			}
			return
		}
		if node.Left != nil {
			dfs(node.Left, path)
		}
		if node.Right != nil {
			dfs(node.Right, path)
		}
	}
	dfs(root, []byte{})
	return result
}

func reverse(input []byte) []byte {
	output := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[len(input)-i-1]
	}
	return output
}
