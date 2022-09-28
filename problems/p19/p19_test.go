package p19

import (
	"github.com/maomaozgw/leetcode/structures/listnode"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				head: listnode.NewG(1, 2, 3, 4, 5),
				n:    2,
			},
			want: listnode.NewG(1, 2, 3, 5),
		},
		{
			name: "Example 2",
			args: args{
				head: listnode.NewG(1),
				n:    1,
			},
			want: listnode.NewG[int](),
		},
		{
			name: "Example 3",
			args: args{
				head: listnode.NewG(1, 2),
				n:    1,
			},
			want: listnode.NewG(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !listnode.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got.String(), tt.want.String())
			}
		})
	}
}
