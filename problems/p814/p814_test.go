package p814

import (
	"github.com/maomaozgw/leetcode/structures/tools"
	"github.com/maomaozgw/leetcode/structures/treenode"
	"testing"
)

func Test_pruneTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				root: treenode.NewBinaryTree(tools.Ptr(1), nil, tools.Ptr(0), tools.Ptr(0), tools.Ptr(1)),
			},
			want: treenode.NewBinaryTree(tools.Ptr(1), nil, tools.Ptr(0), nil, tools.Ptr(1)),
		},
		{
			name: "Example 2",
			args: args{
				root: treenode.NewFullBinaryTree(1, 1, 0, 1, 1, 0, 1, 0),
			},
			want: treenode.NewBinaryTree(tools.Ptr(1), tools.Ptr(1), tools.Ptr(0), tools.Ptr(1), tools.Ptr(1), nil, tools.Ptr(1)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pruneTree(tt.args.root); !treenode.Equal(got, tt.want) {
				t.Errorf("pruneTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
