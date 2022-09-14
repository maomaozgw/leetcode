package p1457

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
	"testing"
)

func Test_pseudoPalindromicPaths(t *testing.T) {
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
				root: treenode.NewBinaryTreeFromString(treenode.MustInt, "2", "3", "1", "3", "1", "null", "1"),
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				root: treenode.NewBinaryTreeFromString(treenode.MustInt, "2", "1", "1", "1", "3", "null", "null", "null", "null", "null", "1"),
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				root: treenode.NewFullBinaryTree(9),
			},
			want: 1,
		},
		{
			name: "Example 4",
			args: args{
				root: treenode.NewBinaryTreeFromString(treenode.MustInt, "2", "1", "1", "3", "2", "null", "1"),
			},
			want: 2,
		},
		{
			name: "WA 1",
			args: args{
				root: treenode.NewBinaryTreeFromString(treenode.MustInt, "9", "5", "4", "5", "null", "2", "6", "2", "5", "null", "8", "3", "9", "2", "3", "1", "1", "null", "4", "5", "4", "2", "2", "6", "4", "null", "null", "1", "7", "null", "5", "4", "7", "null", "null", "7", "null", "1", "5", "6", "1", "null", "null", "null", "null", "9", "2", "null", "9", "7", "2", "1", "null", "null", "null", "6", "null", "null", "null", "null", "null", "null", "null", "null", "null", "5", "null", "null", "3", "null", "null", "null", "8", "null", "1", "null", "null", "8", "null", "null", "null", "null", "2", "null", "8", "7"),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pseudoPalindromicPaths(tt.args.root); got != tt.want {
				t.Errorf("pseudoPalindromicPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
