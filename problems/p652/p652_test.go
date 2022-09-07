package p652

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
	"reflect"
	"testing"
)

func Test_findDuplicateSubtrees(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				root: treenode.NewBinaryTreeFromString(treenode.MustInt, "1", "2", "3", "4", "null", "2", "4", "null", "null", "4"),
			},
			want: []*TreeNode{
				treenode.NewFullBinaryTree(2, 4),
				treenode.NewFullBinaryTree(4),
			},
		},
		{
			name: "Example 2",
			args: args{
				root: treenode.NewBinaryTreeFromString(treenode.MustInt, "2", "1", "1"),
			},
			want: []*TreeNode{
				treenode.NewFullBinaryTree(1),
			},
		},
		{
			name: "WA 1",
			args: args{
				root: treenode.NewBinaryTreeFromString(treenode.MustInt, "0", "0", "0", "0", "null", "null", "0", "null", "null", "null", "0"),
			},
			want: []*TreeNode{
				treenode.NewFullBinaryTree(0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicateSubtrees(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicateSubtrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
