package p113

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
	"reflect"
	"testing"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				root:      treenode.NewBinaryTreeFromString(treenode.MustInt, "5", "4", "8", "11", "null", "13", "4", "7", "2", "null", "null", "5", "1"),
				targetSum: 22,
			},
			want: [][]int{
				{5, 4, 11, 2},
				{5, 8, 4, 5},
			},
		},
		{
			name: "Example 2",
			args: args{
				root:      treenode.NewFullBinaryTree(1, 2, 3),
				targetSum: 5,
			},
			want: [][]int{},
		},
		{
			name: "Example 3",
			args: args{
				root:      treenode.NewFullBinaryTree(1, 2),
				targetSum: 0,
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.targetSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
