package p106

import (
	"testing"

	"github.com/maomaozgw/leetcode/structures/treenode"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				inorder:   []int{9, 3, 15, 20, 7},
				postorder: []int{9, 15, 7, 20, 3},
			},
			want: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `3,9,20,null,null,15,7`),
		},
		{
			name: "Example 2",
			args: args{
				inorder:   []int{-1},
				postorder: []int{-1},
			},
			want: treenode.NewFullBinaryTree[int](-1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.inorder, tt.args.postorder); !treenode.Equal(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
