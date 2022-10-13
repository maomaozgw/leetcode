package p237

import (
	"github.com/maomaozgw/leetcode/structures/listnode"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		root *ListNode
		node *ListNode
	}
	e1 := listnode.NewG(4, 5, 1, 9)
	n5 := e1.Next
	e2 := listnode.NewG(4, 5, 1, 9)
	n1 := e2.Next.Next
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				root: e1,
				node: n5,
			},
			want: listnode.NewG(4, 1, 9),
		},
		{
			name: "Example 2",
			args: args{
				root: e2,
				node: n1,
			},
			want: listnode.NewG(4, 5, 9),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteNode(tt.args.node)
			if !listnode.DeepEqual(tt.args.root, tt.want) {
				t.Errorf("list not equals, got = %s, want %s", tt.args.root.String(), tt.want.String())
			}
		})
	}
}
