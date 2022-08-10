package p109

import (
	"github.com/maomaozgw/leetcode/structures/listnode"
	"github.com/maomaozgw/leetcode/structures/tools"
	"github.com/maomaozgw/leetcode/structures/treenode"
	"reflect"
	"testing"
)

func Test_sortedListToBST(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "",
			args: args{
				head: listnode.NewG(-10, -3, 0, 5, 9),
			},
			want: treenode.NewBinaryTree(tools.Ptr(0), tools.Ptr(-3), tools.Ptr(9), tools.Ptr(-10), nil, tools.Ptr(5)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedListToBST(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedListToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
