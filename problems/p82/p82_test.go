package p82

import (
	"github.com/maomaozgw/leetcode/structures/listnode"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
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
				head: listnode.NewG(1, 2, 3, 3, 4, 4, 5),
			},
			want: listnode.NewG(1, 2, 5),
		},
		{
			name: "Example 2",
			args: args{
				head: listnode.NewG(1, 1, 1, 2, 3),
			},
			want: listnode.NewG(2, 3),
		},
		{
			name: "UE 1",
			args: args{
				head: listnode.NewG(1),
			},
			want: listnode.NewG(1),
		},
		{
			name: "UE 2",
			args: args{
				head: listnode.NewG(1, 1, 1),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !listnode.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
