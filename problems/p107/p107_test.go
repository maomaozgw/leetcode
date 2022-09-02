package p107

import (
	"github.com/maomaozgw/leetcode/structures/tools"
	"github.com/maomaozgw/leetcode/structures/treenode"
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
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
				{15, 7}, {9, 20}, {3},
			},
		},
		{
			name: "Example 2",
			args: args{
				root: treenode.NewBinaryTree(tools.Ptr(1)),
			},
			want: [][]int{{1}},
		},
		{
			name: "Example 3",
			args: args{
				root: nil,
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderBottom(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
