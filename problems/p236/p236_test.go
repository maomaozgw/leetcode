package p236

import (
	"github.com/maomaozgw/leetcode/structures/tools"
	"github.com/maomaozgw/leetcode/structures/treenode"
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	e1Root := treenode.NewBinaryTree(tools.Ptr(3), tools.Ptr(5), tools.Ptr(1), tools.Ptr(6), tools.Ptr(2), tools.Ptr(0), tools.Ptr(8), nil, nil, tools.Ptr(7), tools.Ptr(4))
	tree2 := treenode.NewBinaryTree(tools.Ptr(1), tools.Ptr(2))
	treeW1 := treenode.NewBinaryTree(tools.Ptr(-1), tools.Ptr(0), tools.Ptr(3), tools.Ptr(-2), tools.Ptr(4), nil, nil, tools.Ptr(8))
	treeW2 := treenode.NewBinaryTree(tools.Ptr(37), tools.Ptr(-34), tools.Ptr(-48), nil, tools.Ptr(-100), tools.Ptr(-101), tools.Ptr(48), nil, nil, nil, nil, tools.Ptr(-54), nil, tools.Ptr(-71), tools.Ptr(-22), nil, nil, nil, tools.Ptr(8))
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				root: e1Root,
				p:    treenode.MustFirstByValue(e1Root, 5),
				q:    treenode.MustFirstByValue(e1Root, 1),
			},
			want: treenode.MustFirstByValue(e1Root, 3),
		},
		{
			name: "Example 2",
			args: args{
				root: e1Root,
				p:    treenode.MustFirstByValue(e1Root, 5),
				q:    treenode.MustFirstByValue(e1Root, 4),
			},
			want: treenode.MustFirstByValue(e1Root, 5),
		},
		{
			name: "Example 3",
			args: args{
				root: tree2,
				p:    treenode.MustFirstByValue(tree2, 1),
				q:    treenode.MustFirstByValue(tree2, 2),
			},
			want: treenode.MustFirstByValue(tree2, 1),
		},
		{
			name: "Example 4",
			args: args{
				root: e1Root,
				p:    treenode.MustFirstByValue(e1Root, 6),
				q:    treenode.MustFirstByValue(e1Root, 0),
			},
			want: treenode.MustFirstByValue(e1Root, 3),
		},
		{
			name: "WA 1",
			args: args{
				root: treeW1,
				p:    treenode.MustFirstByValue(treeW1, 3),
				q:    treenode.MustFirstByValue(treeW1, 8),
			},
			want: treenode.MustFirstByValue(treeW1, -1),
		},
		{
			name: "WA 2",
			args: args{
				root: treeW2,
				p:    treenode.MustFirstByValue(treeW2, -71),
				q:    treenode.MustFirstByValue(treeW2, 48),
			},
			want: treenode.MustFirstByValue(treeW2, 48),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
