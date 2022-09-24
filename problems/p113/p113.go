// https://leetcode.com/problems/path-sum-ii/

package p113

import (
	"github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]

func pathSum(root *TreeNode, targetSum int) [][]int {
	stack := linkedliststack.New()
	var result = make([][]int, 0)
	dfs(root, stack, &result, targetSum)
	return result
}

func dfs(root *TreeNode, stack *linkedliststack.Stack, result *[][]int, remainSum int) {
	if root == nil {
		return
	}
	remainSum -= root.Val
	stack.Push(root.Val)
	if root.Left == nil && root.Right == nil {
		// leaf node
		if remainSum == 0 {
			values := stack.Values()
			l := stack.Size()
			var ans = make([]int, l)
			l -= 1
			for i, item := range values {
				ans[l-i] = item.(int)
			}
			*result = append(*result, ans)
		}
		stack.Pop()
		return
	}
	dfs(root.Left, stack, result, remainSum)
	dfs(root.Right, stack, result, remainSum)
	stack.Pop()
}
