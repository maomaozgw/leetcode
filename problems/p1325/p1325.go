package p1325

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	removeLeafLeftNode(root, root.Left, target)
	removeLeafRightNode(root, root.Right, target)
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}

func removeLeafLeftNode(p *TreeNode, n *TreeNode, target int) {
	if n == nil {
		return
	}
	removeLeafLeftNode(n, n.Left, target)
	removeLeafRightNode(n, n.Right, target)
	if n.Left == nil && n.Right == nil && n.Val == target {
		p.Left = nil
	}
}

func removeLeafRightNode(p *TreeNode, n *TreeNode, target int) {
	if n == nil {
		return
	}
	removeLeafLeftNode(n, n.Left, target)
	removeLeafRightNode(n, n.Right, target)
	if n.Left == nil && n.Right == nil && n.Val == target {
		p.Right = nil
	}
}
