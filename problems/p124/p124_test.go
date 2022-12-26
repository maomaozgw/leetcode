package p124

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
	"testing"
)

func Test_maxPathSum(t *testing.T) {
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
			args: args{root: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, "1,2,3")},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{root: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, "-10,9,20,null,null,15,7")},
			want: 42,
		},
		{
			name: "Example 3",
			args: args{root: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, "-3")},
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
