// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

package p235

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p == nil || q == nil {
		return nil
	}
	if p == root || q == root {
		return root
	}
	if p.Val < root.Val && q.Val > root.Val {
		return root
	} else if p.Val > root.Val && q.Val < root.Val {
		return root
	} else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return lowestCommonAncestor(root.Left, p, q)
	}

}
