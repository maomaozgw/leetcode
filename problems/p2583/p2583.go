package p2583

import (
	"container/heap"

	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]

type heapArr struct {
	data []int64
}

func (h *heapArr) Len() int {
	return len(h.data)
}
func (h *heapArr) Less(i, j int) bool {
	return h.data[i] > h.data[j]
}
func (h *heapArr) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
func (h *heapArr) Push(x interface{}) {
	h.data = append(h.data, int64(x.(int)))
}
func (h *heapArr) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	if root == nil {
		return -1
	}
	var q []*TreeNode
	q = append(q, root)
	var levelSum = &heapArr{
		data: []int64{},
	}
	for len(q) > 0 {
		var newQ []*TreeNode
		var levelSumTmp = 0
		for i := 0; i < len(q); i++ {
			c := q[i]
			levelSumTmp += c.Val
			if c.Left != nil {
				newQ = append(newQ, c.Left)
			}
			if c.Right != nil {
				newQ = append(newQ, c.Right)
			}
		}
		heap.Push(levelSum, levelSumTmp)
		q = newQ
	}
	if levelSum.Len() < k {
		return -1
	}
	var v any
	for i := 0; i < k; i++ {
		v = heap.Pop(levelSum)
	}
	return v.(int64)
}
