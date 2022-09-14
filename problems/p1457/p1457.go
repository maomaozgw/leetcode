// https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/

package p1457

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]

func pseudoPalindromicPaths(root *TreeNode) int {
	var totalCount = 0
	m := make([]int, 10)
	dfs(root, m, &totalCount)
	return totalCount
}

func dfs(root *TreeNode, m []int, totalCount *int) {
	if root == nil {
		return
	}
	m[root.Val]++

	dfs(root.Left, m, totalCount)
	dfs(root.Right, m, totalCount)
	if root.Left == nil && root.Right == nil {
		if isPalindrome(m) {
			*totalCount++
		}
	}
	m[root.Val]--
}

func isPalindrome(a []int) bool {
	singleCount := 0
	totalCount := 0
	for i := range a {
		totalCount += a[i]
		if a[i]%2 == 1 {
			singleCount++
			if singleCount > 1 {
				return false
			}
		}
	}
	return singleCount == totalCount%2
}
