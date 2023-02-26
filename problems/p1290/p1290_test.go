package p1290

import (
	"testing"

	"github.com/maomaozgw/leetcode/structures/listnode"
)

func Test_getDecimalValue(t *testing.T) {
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
				head: listnode.NewG[int](1, 0, 1),
			},
			want: 5,
		},
		{
			name: "Example 2",
			args: args{
				head: listnode.NewG[int](0),
			},
			want: 0,
		},
		{
			name: "UE 1",
			args: args{
				head: listnode.NewG[int](1, 0, 0),
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDecimalValue(tt.args.head); got != tt.want {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
