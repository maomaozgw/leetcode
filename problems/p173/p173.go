package p173

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]

type BSTIterator struct {
	q []int
}

func Constructor(root *TreeNode) BSTIterator {
	b := BSTIterator{}
	b.q = []int{}
	dfs(root, &b.q)
	return b
}

func dfs(root *TreeNode, q *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, q)
	*q = append(*q, root.Val)
	dfs(root.Right, q)
}

func (this *BSTIterator) Next() int {
	val := this.q[0]
	this.q = this.q[1:]
	return val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.q) > 0
}
