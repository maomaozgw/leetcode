package p110

import (
	"testing"

	"github.com/maomaozgw/leetcode/structures/treenode"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, `3,9,20,null,null,15,7`),
			},
			want: true,
		},
		{
			name: "Example 1",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, `1,2,2,3,3,null,null,4,4`),
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				root: nil,
			},
			want: true,
		},
		{
			name: "WA 1",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, `1,2,2,3,null,null,3,4,null,null,4`),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
