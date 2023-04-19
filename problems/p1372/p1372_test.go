package p1372

import (
	"testing"

	"github.com/maomaozgw/leetcode/structures/treenode"
)

func Test_longestZigZag(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1 ",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1`),
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `1,1,1,null,1,null,null,1,1,null,1`),
			},
			want: 4,
		},
		{
			name: "Example 3",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `1`),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestZigZag(tt.args.root); got != tt.want {
				t.Errorf("longestZigZag() = %v, want %v", got, tt.want)
			}
		})
	}
}
