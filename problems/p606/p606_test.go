package p606

import (
	"github.com/maomaozgw/leetcode/structures/tools"
	"github.com/maomaozgw/leetcode/structures/treenode"
	"testing"
)

func Test_tree2str(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				root: treenode.NewFullBinaryTree(1, 2, 3, 4),
			},
			want: "1(2(4))(3)",
		},
		{
			name: "Example 2",
			args: args{
				root: treenode.NewBinaryTree(tools.Ptr(1), tools.Ptr(2), tools.Ptr(3), nil, tools.Ptr(4)),
			},
			want: "1(2()(4))(3)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree2str(tt.args.root); got != tt.want {
				t.Errorf("tree2str() = %v, want %v", got, tt.want)
			}
		})
	}
}
