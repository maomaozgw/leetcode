package p112

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				root:      treenode.NewBinaryTreeFromString(treenode.MustInt, "5", "4", "8", "11", "null", "13", "4", "7", "2", "null", "null", "null", "1"),
				targetSum: 22,
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				root:      treenode.NewFullBinaryTree(1, 2, 3),
				targetSum: 5,
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				root:      nil,
				targetSum: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
