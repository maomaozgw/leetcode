// https://leetcode.com/problems/count-complete-tree-nodes/

package p222

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
	"math"
)

type TreeNode = treenode.BinaryTree[int]

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := getDepth(root)
	if depth == 0 {
		return 1
	}
	maxNodes := int(math.Pow(2, float64(depth)))
	left, right, last := 0, maxNodes-1, 0
	for left <= right {
		mid := (right-left)/2 + left
		if binarySearch(mid, maxNodes, depth, root) {
			last = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return maxNodes + last
}

func getDepth(root *TreeNode) int {
	var depth = 0
	var current = root
	for current.Left != nil {
		current = current.Left
		depth++
	}
	return depth
}

func binarySearch(idx int, max, depth int, node *TreeNode) bool {
	left, right := 0, max-1
	for i := 0; i < depth; i++ {
		mid := (right-left)/2 + left
		if idx <= mid {
			if node.Left == nil {
				return false
			}
			right = mid - 1
			node = node.Left
			continue
		}
		if node.Right == nil {
			return false
		}
		left = mid + 1
		node = node.Right
	}
	return node != nil
}
