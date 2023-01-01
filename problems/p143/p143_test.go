package p143

import (
	"github.com/maomaozgw/leetcode/structures/listnode"
	"testing"
)

func Test_reorderList(t *testing.T) {
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
				head: listnode.NewG(1, 2, 3, 4),
			},
			want: listnode.NewG(1, 4, 2, 3),
		},
		{
			name: "Example 2",
			args: args{
				head: listnode.NewG(1, 2, 3, 4, 5),
			},
			want: listnode.NewG(1, 5, 2, 4, 3),
		},
		{
			name: "UE 1",
			args: args{
				head: nil,
			},
			want: nil,
		},
		{
			name: "UE 2",
			args: args{
				head: listnode.NewG(1),
			},
			want: listnode.NewG(1),
		},
		{
			name: "UE 3",
			args: args{
				head: listnode.NewG(1, 2),
			},
			want: listnode.NewG(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
			if !listnode.DeepEqual(tt.args.head, tt.want) {
				t.Errorf("reorderList failed: want %s got %s", tt.want, tt.args.head)
			}
		})
	}
}
