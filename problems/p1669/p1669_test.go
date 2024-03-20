package p1669

import (
	"reflect"
	"testing"

	"github.com/maomaozgw/leetcode/structures/listnode"
)

func Test_mergeInBetween(t *testing.T) {
	type args struct {
		list1 *ListNode
		a     int
		b     int
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				list1: listnode.NewG[int](10, 1, 13, 6, 9, 5),
				a:     3,
				b:     4,
				list2: listnode.NewG[int](1000000, 1000001, 1000002),
			},
			want: listnode.NewG[int](10, 1, 13, 1000000, 1000001, 1000002, 5),
		},
		{
			name: "Example 2",
			args: args{
				list1: listnode.NewG[int](0, 1, 2, 3, 4, 5, 6),
				a:     2,
				b:     5,
				list2: listnode.NewG[int](1000000, 1000001, 1000002, 1000003, 1000004),
			},
			want: listnode.NewG[int](0, 1, 1000000, 1000001, 1000002, 1000003, 1000004, 6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeInBetween(tt.args.list1, tt.args.a, tt.args.b, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeInBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
