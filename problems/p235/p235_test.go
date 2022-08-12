package p235

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
	var p = tools.Ptr[int]
	e1 := treenode.NewBinaryTree(p(6), p(2), p(8), p(0), p(4), p(7), p(9), nil, nil, p(3), p(5))
	e2 := treenode.NewFullBinaryTree(2, 1)
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				root: e1,
				p:    treenode.MustFirstByValue(e1, 2),
				q:    treenode.MustFirstByValue(e1, 4),
			},
			want: treenode.MustFirstByValue(e1, 2),
		},
		{
			name: "Example 2",
			args: args{
				root: e2,
				p:    treenode.MustFirstByValue(e2, 2),
				q:    treenode.MustFirstByValue(e2, 1),
			},
			want: treenode.MustFirstByValue(e2, 2),
		},
		{
			name: "Example 3",
			args: args{
				root: e1,
				p:    treenode.MustFirstByValue(e1, 2),
				q:    treenode.MustFirstByValue(e1, 8),
			},
			want: treenode.MustFirstByValue(e1, 6),
		},
		{
			name: "Example 3",
			args: args{
				root: e1,
				p:    treenode.MustFirstByValue(e1, 2),
				q:    treenode.MustFirstByValue(e1, 8),
			},
			want: treenode.MustFirstByValue(e1, 6),
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
