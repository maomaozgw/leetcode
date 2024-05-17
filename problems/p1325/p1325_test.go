package p1325

import (
	"reflect"
	"testing"

	"github.com/maomaozgw/leetcode/structures/treenode"
)

func Test_removeLeafNodes(t *testing.T) {
	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				root:   treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `1,2,3,2,null,2,4`),
				target: 2,
			},
			want: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `1,null,3,null,4`),
		},
		{
			name: "Example 2",
			args: args{
				root:   treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `1,3,3,3,2`),
				target: 3,
			},
			want: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `1,3,null,null,2`),
		},
		{
			name: "Example 3",
			args: args{
				root:   treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `1,2,null,2,null,2`),
				target: 2,
			},
			want: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `1`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeLeafNodes(tt.args.root, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeLeafNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
