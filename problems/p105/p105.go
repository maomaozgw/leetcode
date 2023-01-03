package p105

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, inorder, 0, 0, len(inorder)-1)
}

func build(preorder []int, inorder []int, preIdx, inStart, inEnd int) *TreeNode {
	if inStart < 0 || inStart == len(inorder) || inStart > inEnd {
		return nil
	}
	rootVal := preorder[preIdx]
	root := &TreeNode{
		Val: rootVal,
	}
	idx := inStart
	for ; idx < len(inorder); idx++ {
		if inorder[idx] == rootVal {
			// 左侧根节点就在当前根节点的下一个
			root.Left = build(preorder, inorder, preIdx+1, inStart, idx-1)
			// 右侧根节点需要跳过左侧的节点个数
			root.Right = build(preorder, inorder, preIdx+(idx-inStart+1), idx+1, inEnd)
			break
		}
	}
	return root
}
