package p2130

import (
	"testing"

	"github.com/maomaozgw/leetcode/structures/listnode"
)

func Test_pairSum(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				head: listnode.NewG(5, 4, 2, 1),
			},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{
				head: listnode.NewG(4, 2, 2, 3),
			},
			want: 7,
		},
		{
			name: "Example 3",
			args: args{
				head: listnode.NewG(1, 100000),
			},
			want: 100001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pairSum(tt.args.head); got != tt.want {
				t.Errorf("pairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
