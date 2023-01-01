package p25

import (
	"github.com/maomaozgw/leetcode/structures/listnode"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
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
				k:    2,
			},
			want: listnode.NewG(2, 1, 4, 3, 5),
		},
		{
			name: "Example 2",
			args: args{
				head: listnode.NewG(1, 2, 3, 4, 5),
				k:    3,
			},
			want: listnode.NewG(3, 2, 1, 4, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !listnode.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
