package p1339

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
	"testing"
)

func Test_maxProduct(t *testing.T) {
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
			args: args{root: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, "1,2,3,4,5,6")},
			want: 110,
		},
		{
			name: "Example 2",
			args: args{root: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, "1,null,2,3,4,null,null,5,6")},
			want: 90,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.root); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
