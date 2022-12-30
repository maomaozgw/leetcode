package p142

import (
	"github.com/maomaozgw/leetcode/structures/listnode"
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	l1 := listnode.NewG(3, 2, 0, -4)
	e2 := l1.Lookup(2)
	e4 := l1.Lookup(-4)
	e4.Next = e2
	l2 := listnode.NewG(1, 2)
	l2e1 := l2.Lookup(1)
	l2e2 := l2.Lookup(2)
	l2e2.Next = l2e1
	l3 := listnode.NewG(1)
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				head: l1,
			},
			want: e2,
		},
		{
			name: "Example 2",
			args: args{
				head: l2,
			},
			want: l2e1,
		},
		{
			name: "Example 3",
			args: args{
				head: l3,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got.Val, tt.want.Val)
			}
		})
	}
}
