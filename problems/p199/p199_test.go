package p199

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, `1,2,3,null,5,null,4`),
			},
			want: []int{1, 3, 4},
		},
		{
			name: "Example 2",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, `1,null,3`),
			},
			want: []int{1, 3},
		},
		{
			name: "Example 3",
			args: args{
				root: nil,
			},
			want: []int{},
		},
		{
			name: "WA 1",
			args: args{
				root: treenode.NewFullBinaryTree(1, 2, 3, 4),
			},
			want: []int{1, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
