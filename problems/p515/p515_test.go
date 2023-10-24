package p515

import (
	"reflect"
	"testing"

	"github.com/maomaozgw/leetcode/structures/treenode"
)

func Test_largestValues(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `1,3,2,5,3,null,9`),
			},
			want: []int{1, 3, 9},
		},
		{
			name: "Example 2",
			args: args{
				root: treenode.NewBinaryTreeFromLeetCodeStr[int](treenode.MustInt, `1,2,3`),
			},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestValues(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
