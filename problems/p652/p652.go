// https://leetcode.com/problems/find-duplicate-subtrees/

package p652

import (
	"fmt"
	"github.com/maomaozgw/leetcode/structures/treenode"
	"strconv"
)

type TreeNode = treenode.BinaryTree[int]

const (
	reported = -100
)

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var result []*TreeNode
	var m = map[string]int{}
	dfs(root, m, &result)
	return result
}

func dfs(root *TreeNode, m map[string]int, result *[]*TreeNode) string {
	if root == nil {
		return "-"
	}
	var l, f = dfs(root.Left, m, result), dfs(root.Right, m, result)
	var key = l + f + "-" + strconv.Itoa(root.Val)
	if m[key] == 0 {
		m[key]++
	} else if m[key] == 1 {
		fmt.Println(key)
		*result = append(*result, root)
		m[key] = reported
	}
	return key
}
