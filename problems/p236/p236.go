// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

package p236

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]

const (
	findNull = 0
	findP    = 1
	findQ    = 2
	findAll  = 3
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q {
		return root
	}
	result, _ := bfFindInTree(root, p, q)
	return result
}

func bfFindOneInTree(root, one *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Left == one || root.Right == one {
		return true
	}
	if bfFindOneInTree(root.Left, one) {
		return true
	}
	if bfFindOneInTree(root.Right, one) {
		return true
	}
	return false
}

func bfFindInTree(root, p, q *TreeNode) (*TreeNode, int) {
	if root == nil {
		return nil, findNull
	}
	if root.Left == p && root.Right == q {
		return root, findAll
	}
	if root.Right == p && root.Left == q {
		return root, findAll
	}
	if root.Left == p || root.Left == q {
		target := p
		find := findQ
		if root.Left == p {
			find = findP
			target = q
		}

		ok := bfFindOneInTree(root.Left, target)
		if ok {
			return root.Left, findAll
		}
		ok = bfFindOneInTree(root.Right, target)
		if ok {
			return root, findAll
		}
		return nil, find
	} else if root.Right == p || root.Right == q {
		target := p
		find := findQ

		if root.Right == p {
			find = findP
			target = q
		}
		ok := bfFindOneInTree(root.Right, target)
		if ok {
			return root.Right, findAll
		}
		ok = bfFindOneInTree(root.Left, target)
		if ok {
			return root, findAll
		}
		return nil, find
	}
	parent, findResult := bfFindInTree(root.Left, p, q)
	if findResult == findAll {
		return parent, findResult
	} else if findResult == findP {
		if bfFindOneInTree(root.Right, q) {
			return root, findAll
		}
		return nil, findResult
	} else if findResult == findQ {
		if bfFindOneInTree(root.Right, p) {
			return root, findAll
		}
		return nil, findResult
	}
	return bfFindInTree(root.Right, p, q)
}
