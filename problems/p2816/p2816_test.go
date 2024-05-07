package p2816

import (
	"testing"

	"github.com/maomaozgw/leetcode/structures/listnode"
)

func Test_doubleIt(t *testing.T) {
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
				head: listnode.NewG[int](1, 8, 9),
			},
			want: listnode.NewG[int](3, 7, 8),
		},
		{
			name: "Example 2",
			args: args{
				head: listnode.NewG[int](9, 9, 9),
			},
			want: listnode.NewG[int](1, 9, 9, 8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doubleIt(tt.args.head); !listnode.DeepEqual[int](got, tt.want) {
				t.Errorf("doubleIt() = %v, want %v", got, tt.want)
			}
		})
	}
}
