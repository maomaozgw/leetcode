package p1448

import (
	"github.com/maomaozgw/leetcode/structures/tools"
	"github.com/maomaozgw/leetcode/structures/treenode"
	"testing"
)

func Test_goodNodes(t *testing.T) {
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
				root: treenode.NewBinaryTree(tools.Ptr(3), tools.Ptr(1), tools.Ptr(4), tools.Ptr(3), nil, tools.Ptr(1), tools.Ptr(5)),
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				root: treenode.NewBinaryTree(tools.Ptr(3), tools.Ptr(3), nil, tools.Ptr(4), tools.Ptr(2)),
			},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{
				root: treenode.NewBinaryTree(tools.Ptr(1)),
			},
			want: 1,
		},
		{
			name: "WA 1",
			args: args{
				root: treenode.NewBinaryTree(tools.Ptr(2), tools.Ptr(4), tools.Ptr(4), tools.Ptr(4), nil, tools.Ptr(1), tools.Ptr(3), nil, nil, tools.Ptr(5), nil, nil, nil, nil, tools.Ptr(5), tools.Ptr(4), tools.Ptr(4)),
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goodNodes(tt.args.root); got != tt.want {
				t.Errorf("goodNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
