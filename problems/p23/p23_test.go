package p23

import (
	"reflect"
	"testing"

	"github.com/maomaozgw/leetcode/structures/listnode"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				lists: []*ListNode{
					listnode.NewG(1, 4, 5),
					listnode.NewG(1, 3, 4),
					listnode.NewG(2, 6),
				},
			},
			want: listnode.NewG(1, 1, 2, 3, 4, 4, 5, 6),
		},
		{
			name: "Example 2",
			args: args{
				lists: []*ListNode{},
			},
			want: nil,
		},
		{
			name: "Example 3",
			args: args{
				lists: []*ListNode{nil},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
