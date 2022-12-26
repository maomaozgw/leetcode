package p328

import (
	"github.com/maomaozgw/leetcode/structures/listnode"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
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
				head: listnode.NewG(1, 2, 3, 4, 5),
			},
			want: listnode.NewG(1, 3, 5, 2, 4),
		},
		{
			name: "Example 2",
			args: args{
				head: listnode.NewG(2, 1, 3, 5, 6, 4, 7),
			},
			want: listnode.NewG(2, 3, 6, 7, 1, 5, 4),
		},
		{
			name: "Example 3",
			args: args{
				head: listnode.NewG(2, 1, 3, 5, 6, 4),
			},
			want: listnode.NewG(2, 3, 6, 1, 5, 4),
		},
		{
			name: "Example 4",
			args: args{
				head: listnode.NewG(1),
			},
			want: listnode.NewG(1),
		},
		{
			name: "Example 5",
			args: args{
				head: listnode.NewG(1, 2),
			},
			want: listnode.NewG(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddEvenList(tt.args.head); !listnode.DeepEqual(got, tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", got, tt.want)
			}
		})
	}
}
