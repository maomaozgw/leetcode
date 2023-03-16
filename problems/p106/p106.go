package p106

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]

func buildTree(inorder []int, postorder []int) *TreeNode {
	var inorderMap = map[int]int{}
	for idx := range inorder {
		inorderMap[inorder[idx]] = idx
	}
	tmp := len(postorder) - 1
	return build(postorder, inorderMap, &tmp, 0, len(inorder)-1)
}

func build(postorder []int, inorder map[int]int, postIdx *int, inStart, inEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}
	rootVal := postorder[*postIdx]
	*postIdx--
	root := &TreeNode{
		Val: rootVal,
	}
	idx := inorder[rootVal]
	root.Right = build(postorder, inorder, postIdx, idx+1, inEnd)
	root.Left = build(postorder, inorder, postIdx, inStart, idx-1)
	return root
}
