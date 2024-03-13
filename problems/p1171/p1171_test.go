package p1171

import (
	"reflect"
	"testing"

	"github.com/maomaozgw/leetcode/structures/listnode"
)

func Test_removeZeroSumSublists(t *testing.T) {
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
				head: listnode.NewG[int](1, 2, -3, 3, 1),
			},
			want: listnode.NewG[int](3, 1),
		},
		{
			name: "Example 2",
			args: args{
				head: listnode.NewG[int](1, 2, 3, -3, 4),
			},
			want: listnode.NewG[int](1, 2, 4),
		},
		{
			name: "Example 3",
			args: args{
				head: listnode.NewG[int](1, 2, 3, -3, -2),
			},
			want: listnode.NewG[int](1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeZeroSumSublists(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeZeroSumSublists() = %v, want %v", got, tt.want)
			}
		})
	}
}
