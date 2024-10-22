package p2583

import (
	"testing"

	"github.com/maomaozgw/leetcode/structures/treenode"
)

func Test_kthLargestLevelSum(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `5,8,9,2,1,3,7,4,6`),
				k:    2,
			},
			want: 13,
		},
		{
			name: "Example 2",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `1,2,null,3`),
				k:    1,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargestLevelSum(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthLargestLevelSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
