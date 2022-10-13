package p2095

import (
	"github.com/maomaozgw/leetcode/structures/listnode"
	"testing"
)

func Test_deleteMiddle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				head: listnode.NewG(1, 3, 4, 7, 1, 2, 6),
			},
			want: listnode.NewG(1, 3, 4, 1, 2, 6),
		},
		{
			name: "Example 2",
			args: args{
				head: listnode.NewG(1, 2, 3, 4),
			},
			want: listnode.NewG(1, 2, 4),
		},
		{
			name: "Example 3",
			args: args{
				head: listnode.NewG(2, 1),
			},
			want: listnode.NewG(2),
		},
		{
			name: "P1",
			args: args{
				head: listnode.NewG(1),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteMiddle(tt.args.head); !listnode.DeepEqual(got, tt.want) {
				t.Errorf("deleteMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}
