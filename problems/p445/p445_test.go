package p445

import (
	"reflect"
	"testing"

	"github.com/maomaozgw/leetcode/structures/listnode"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				l1: listnode.NewG(7, 2, 4, 3),
				l2: listnode.NewG(5, 6, 4),
			},
			want: listnode.NewG(7, 8, 0, 7),
		},
		{
			name: "Example 2",
			args: args{
				l1: listnode.NewG(2, 4, 3),
				l2: listnode.NewG(5, 6, 4),
			},
			want: listnode.NewG(8, 0, 7),
		},
		{
			name: "Example 3",
			args: args{
				l1: listnode.NewG(0),
				l2: listnode.NewG(0),
			},
			want: listnode.NewG(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
