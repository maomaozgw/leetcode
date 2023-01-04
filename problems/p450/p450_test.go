package p450

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		root *TreeNode
		key  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, `5,3,6,2,4,null,7`),
				key:  3,
			},
			want: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, `5,4,6,2,null,null,7`),
		},
		{
			name: "Example 2",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, `5,3,6,2,4,null,7`),
				key:  0,
			},
			want: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, `5,3,6,2,4,null,7`),
		},
		{
			name: "Example 3",
			args: args{
				root: nil,
				key:  0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.root, tt.args.key); !treenode.Equal(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
