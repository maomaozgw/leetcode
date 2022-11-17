package p222

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
	"testing"
)

func Test_countNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, "1,2,3,4,5,6"),
			},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{
				root: nil,
			},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{
				root: treenode.NewFullBinaryTree(1),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodes(tt.args.root); got != tt.want {
				t.Errorf("countNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
