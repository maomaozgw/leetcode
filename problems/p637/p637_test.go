package p637

import (
	"github.com/maomaozgw/leetcode/structures/tools"
	"github.com/maomaozgw/leetcode/structures/treenode"
	"reflect"
	"testing"
)

func Test_averageOfLevels(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "Example 1",
			args: args{
				root: treenode.NewBinaryTree(tools.Ptr(3), tools.Ptr(9), tools.Ptr(20), nil, nil, tools.Ptr(15), tools.Ptr(7)),
			},
			want: []float64{
				3.0, 14.5, 11.0,
			},
		},
		{
			name: "Example 2",
			args: args{
				root: treenode.NewBinaryTree(tools.Ptr(3), tools.Ptr(9), tools.Ptr(20), tools.Ptr(15), tools.Ptr(7)),
			},
			want: []float64{
				3.0, 14.5, 11.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageOfLevels(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("averageOfLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
