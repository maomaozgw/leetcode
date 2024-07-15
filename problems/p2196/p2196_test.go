package p2196

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
	"github.com/maomaozgw/leetcode/structures/treenode"
)

func Test_createBinaryTree(t *testing.T) {
	type args struct {
		descriptions [][]int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				descriptions: tools.NewGridFromStr(strconv.Atoi, `[20,15,1],[20,17,0],[50,20,1],[50,80,0],[80,19,1]`),
			},
			want: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `50,20,80,15,17,19`),
		},
		{
			name: "Example 2",
			args: args{
				descriptions: tools.NewGridFromStr(strconv.Atoi, `[1,2,1],[2,3,0],[3,4,1]`),
			},
			want: treenode.NewBinaryTreeFromLeetCodeStr(treenode.MustInt, `1,2,null,null,3,4`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createBinaryTree(tt.args.descriptions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
