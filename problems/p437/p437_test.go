package p437

import (
	"testing"

	"github.com/maomaozgw/leetcode/structures/treenode"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				root:      treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `10,5,-3,3,2,null,11,3,-2,null,1`),
				targetSum: 8,
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				root:      treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `5,4,8,11,null,13,4,7,2,null,null,5,1`),
				targetSum: 22,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
