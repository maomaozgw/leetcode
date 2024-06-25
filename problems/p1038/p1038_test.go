package p1038

import (
	"reflect"
	"testing"

	"github.com/maomaozgw/leetcode/structures/treenode"
)

func Test_bstToGst(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `4,1,6,0,2,5,7,null,null,null,3,null,null,null,8`),
			},
			want: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, `30,36,21,36,35,26,15,null,null,null,33,null,null,null,8`),
		},
		{
			name: "Example 2",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, `0,null,1`),
			},
			want: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, `1,null,1`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bstToGst(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstToGst() = %v, want %v", got, tt.want)
			}
		})
	}
}
