package p623

import (
	"github.com/maomaozgw/leetcode/structures/tools"
	"github.com/maomaozgw/leetcode/structures/treenode"
	"testing"
)

func Test_addOneRow(t *testing.T) {
	type args struct {
		root  *TreeNode
		val   int
		depth int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				root:  treenode.NewBinaryTree[int](tools.Ptr(4), tools.Ptr(2), tools.Ptr(6), tools.Ptr(3), tools.Ptr(1), tools.Ptr(5)),
				val:   1,
				depth: 2,
			},
			want: treenode.NewBinaryTree(tools.Ptr(4), tools.Ptr(1), tools.Ptr(1), tools.Ptr(2), nil, nil, tools.Ptr(6), tools.Ptr(3), tools.Ptr(1), tools.Ptr(5)),
		},
		{
			name: "Example 2",
			args: args{
				root:  treenode.NewBinaryTree[int](tools.Ptr(4), tools.Ptr(2), nil, tools.Ptr(3), tools.Ptr(1)),
				val:   1,
				depth: 3,
			},
			want: treenode.NewBinaryTree[int](tools.Ptr(4), tools.Ptr(2), nil, tools.Ptr(1), tools.Ptr(1), tools.Ptr(3), nil, nil, tools.Ptr(1)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addOneRow(tt.args.root, tt.args.val, tt.args.depth); !treenode.Equal(got, tt.want) {
				t.Errorf("addOneRow() = %v, want %v", treenode.String(got), treenode.String(tt.want))
			}
		})
	}
}
