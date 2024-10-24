package p951

import (
	"testing"

	"github.com/maomaozgw/leetcode/structures/treenode"
)

func Test_flipEquiv(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				root1: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `1,2,3,4,5,6,null,null,null,7,8`),
				root2: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `1,3,2,null,6,4,5,null,null,null,null,8,7`),
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				root1: nil,
				root2: nil,
			},
			want: true,
		},
		{
			name: "Example 3",
			args: args{
				root1: nil,
				root2: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `1`),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipEquiv(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("flipEquiv() = %v, want %v", got, tt.want)
			}
		})
	}
}
