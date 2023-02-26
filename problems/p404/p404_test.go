package p404

import (
	"testing"

	"github.com/maomaozgw/leetcode/structures/treenode"
)

func Test_sumOfLeftLeaves(t *testing.T) {
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
				root: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `3,9,20,null,null,15,7`),
			},
			want: 24,
		},
		{
			name: "Example 2",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `1`),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfLeftLeaves(tt.args.root); got != tt.want {
				t.Errorf("sumOfLeftLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
}
