package p98

import (
	"github.com/maomaozgw/leetcode/structures/tools"
	"github.com/maomaozgw/leetcode/structures/treenode"
	"testing"
)

func Test_isValidBST(t *testing.T) {
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
				root: treenode.NewBinaryTree(tools.Ptr(2), tools.Ptr(1), tools.Ptr(3)),
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				root: treenode.NewBinaryTree(tools.Ptr(5), tools.Ptr(1), tools.Ptr(4), nil, nil, tools.Ptr(3), tools.Ptr(6)),
			},
			want: false,
		},
		{
			name: "WA 1",
			args: args{
				root: treenode.NewBinaryTree(tools.Ptr(1), tools.Ptr(1)),
			},
			want: false,
		},
		{
			name: "WA 1",
			args: args{
				root: treenode.NewFullBinaryTree(120, 70, 140, 50, 100, 130, 160, 20, 55, 75, 110, 119, 135, 150, 200),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
