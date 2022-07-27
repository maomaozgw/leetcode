package p114

import (
	"github.com/maomaozgw/leetcode/structures/tools"
	"github.com/maomaozgw/leetcode/structures/treenode"
	"testing"
)

func Test_flatten(t *testing.T) {
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
				root: treenode.NewBinaryTree(tools.Ptr(1), tools.Ptr(2), tools.Ptr(5), tools.Ptr(3), tools.Ptr(4), nil, tools.Ptr(6)),
			},
			want: treenode.NewBinaryTree(tools.Ptr(1), nil, tools.Ptr(2), nil, tools.Ptr(3), nil, tools.Ptr(4), nil, tools.Ptr(5), nil, tools.Ptr(6)),
		},
		{
			name: "Example 2",
			args: args{},
		},
		{
			name: "Example 3",
			args: args{},
		},
		{
			name: "Custom 1",
			args: args{
				root: treenode.NewBinaryTree(tools.Ptr(1), tools.Ptr(2), tools.Ptr(5), tools.Ptr(3), tools.Ptr(4), tools.Ptr(6), tools.Ptr(7)),
			},
			want: treenode.NewBinaryTree(tools.Ptr(1), nil, tools.Ptr(2), nil, tools.Ptr(3), nil, tools.Ptr(4), nil, tools.Ptr(5), nil, tools.Ptr(6), nil, tools.Ptr(7)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.args.root)
			if !treenode.Equal(tt.args.root, tt.want) {
				t.Errorf("flatten = %v, want %v", tt.args.root, tt.want)
				treenode.Print(tt.args.root)
				treenode.Print(tt.want)
			}
		})
	}
}
