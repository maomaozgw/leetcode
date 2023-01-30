package p617

import (
	"reflect"
	"testing"

	"github.com/maomaozgw/leetcode/structures/treenode"
)

func Test_mergeTrees(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				root1: treenode.NewFullBinaryTree(1, 3, 2, 5),
				root2: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, `2,1,3,null,4,null,7`),
			},
			want: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, `3,4,5,5,4,null,7`),
		},
		{
			name: "Example 2",
			args: args{
				root1: treenode.NewFullBinaryTree(1),
				root2: treenode.NewFullBinaryTree(1, 2),
			},
			want: treenode.NewFullBinaryTree(2, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTrees(tt.args.root1, tt.args.root2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
