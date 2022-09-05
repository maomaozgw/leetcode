package p429

import (
	"github.com/maomaozgw/leetcode/structures/tools"
	"github.com/maomaozgw/leetcode/structures/treenode"
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				root: treenode.NewNA(tools.Ptr(1), nil, tools.Ptr(3), tools.Ptr(2), tools.Ptr(4), nil, tools.Ptr(5), tools.Ptr(6)),
			},
			want: [][]int{
				{1},
				{3, 2, 4},
				{5, 6},
			},
		},
		{
			name: "Example 2",
			args: args{
				root: treenode.NewNA(
					tools.Ptr(1), nil, tools.Ptr(2), tools.Ptr(3), tools.Ptr(4),
					tools.Ptr(5), nil, nil, tools.Ptr(6), tools.Ptr(7), nil, tools.Ptr(8), nil,
					tools.Ptr(9), tools.Ptr(10), nil, nil, tools.Ptr(11), nil, tools.Ptr(12),
					nil, tools.Ptr(13), nil, nil, tools.Ptr(14),
				),
			},
			want: [][]int{
				{1}, {2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}, {14},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
