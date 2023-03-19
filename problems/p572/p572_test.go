package p572

import (
	"testing"

	"github.com/maomaozgw/leetcode/structures/treenode"
)

func Test_isSubtree(t *testing.T) {
	type args struct {
		root    *TreeNode
		subRoot *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				root:    treenode.NewFullBinaryTree[int](3, 4, 5, 1, 2),
				subRoot: treenode.NewFullBinaryTree[int](4, 1, 2),
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				root:    treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `3,4,5,1,2,null,null,null,null,0`),
				subRoot: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `4,1,2`),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.root, tt.args.subRoot); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
