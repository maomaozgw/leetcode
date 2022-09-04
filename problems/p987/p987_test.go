package p987

import (
	"github.com/maomaozgw/leetcode/structures/tools"
	"github.com/maomaozgw/leetcode/structures/treenode"
	"reflect"
	"testing"
)

func Test_verticalTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				root: treenode.NewBinaryTree(tools.Ptr(3), tools.Ptr(9), tools.Ptr(20), nil, nil, tools.Ptr(15), tools.Ptr(7)),
			},
			want: [][]int{
				{9},
				{3, 15},
				{20},
				{7},
			},
		},
		{
			name: "Example 2",
			args: args{
				root: treenode.NewFullBinaryTree(1, 2, 3, 4, 5, 6, 7),
			},
			want: [][]int{
				{4},
				{2},
				{1, 5, 6},
				{3},
				{7},
			},
		},
		{
			name: "Example 3",
			args: args{
				root: treenode.NewFullBinaryTree(1, 2, 3, 4, 6, 5, 7),
			},
			want: [][]int{
				{4},
				{2},
				{1, 5, 6},
				{3},
				{7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verticalTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verticalTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
