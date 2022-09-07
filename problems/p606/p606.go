// https://leetcode.com/problems/construct-string-from-binary-tree/

package p606

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
	"strconv"
	"strings"
)

type TreeNode = treenode.BinaryTree[int]

const (
	left  = '('
	right = ')'
)

func tree2str(root *TreeNode) string {
	b := &strings.Builder{}
	dfs(root, b)
	return b.String()
}

func dfs(root *TreeNode, builder *strings.Builder) {
	if root == nil {
		return
	}
	builder.WriteString(strconv.Itoa(root.Val))
	if root.Left != nil {
		builder.WriteByte(left)
		dfs(root.Left, builder)
		builder.WriteByte(right)
	}
	if root.Right != nil {
		if root.Left == nil {
			builder.WriteByte(left)
			builder.WriteByte(right)
		}
		builder.WriteByte(left)
		dfs(root.Right, builder)
		builder.WriteByte(right)
	}
}
