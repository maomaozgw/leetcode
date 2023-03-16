package p129

import (
	"testing"

	"github.com/maomaozgw/leetcode/structures/treenode"
)

func Test_sumNumbers(t *testing.T) {
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
				root: treenode.NewFullBinaryTree[int](1, 2, 3),
			},
			want: 25,
		},
		{
			name: "Example 2",
			args: args{
				root: treenode.NewFullBinaryTree[int](4, 9, 0, 5, 1),
			},
			want: 1026,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNumbers(tt.args.root); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
