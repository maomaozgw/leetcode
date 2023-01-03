package p105

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, `3,9,20,null,null,15,7`),
		},
		{
			name: "Example 2",
			args: args{
				preorder: []int{-1},
				inorder:  []int{-1},
			},
			want: treenode.NewFullBinaryTree(-1),
		},
		{
			name: "WA 1",
			args: args{
				preorder: []int{3, 1, 2, 4},
				inorder:  []int{1, 2, 3, 4},
			},
			want: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, `3,1,4,null,2`),
		},
		{
			name: "WA 2",
			args: args{
				preorder: []int{7, -10, -4, 3, -1, 2, -8, 11},
				inorder:  []int{-4, -10, 3, -1, 7, 11, -8, 2},
			},
			want: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, `7,-10,2,-4,3,-8,null,null,null,null,-1,11`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !treenode.Equal(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", treenode.String(got), treenode.String(tt.want))
			}
		})
	}
}
