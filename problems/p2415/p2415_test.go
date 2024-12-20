package p2415

import (
	"testing"

	"github.com/maomaozgw/leetcode/structures/treenode"
)

func Test_reverseOddLevels(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `2,3,5,8,13,21,34`),
			},
			want: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `2,5,3,8,13,21,34`),
		},
		{
			name: "Example 2",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `7,13,11`),
			},
			want: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `7,11,13`),
		},
		{
			name: "Example 3",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `0,1,2,0,0,0,0,1,1,1,1,2,2,2,2`),
			},
			want: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `0,2,1,0,0,0,0,2,2,2,2,1,1,1,1`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseOddLevels(tt.args.root); !treenode.Equal(got, tt.want) {
				t.Errorf("reverseOddLevels() = %v, want %v", treenode.String(got), treenode.String(tt.want))
			}
		})
	}
}
