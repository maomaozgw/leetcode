package p108

import (
	"github.com/maomaozgw/leetcode/structures/tools"
	"github.com/maomaozgw/leetcode/structures/treenode"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{-10, -3, 0, 5, 9},
			},
			want: treenode.NewBinaryTree(tools.Ptr(0), tools.Ptr(-3), tools.Ptr(9), tools.Ptr(-10), nil, tools.Ptr(5)),
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{3, 1},
			},
			want: treenode.NewBinaryTree(tools.Ptr(1), tools.Ptr(3)),
		},
		{
			name: "WA 1",
			args: args{
				nums: []int{3, 5, 8},
			},
			want: treenode.NewBinaryTree(tools.Ptr(5), tools.Ptr(3), tools.Ptr(8)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !treenode.Equal(got, tt.want) {
				t.Errorf("sortedArrayToBST() = \n%v, want\n %v", treenode.String(got), treenode.String(tt.want))
			}
		})
	}
}
