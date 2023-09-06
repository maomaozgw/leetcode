package p725

import (
	"reflect"
	"testing"

	"github.com/maomaozgw/leetcode/structures/listnode"
)

func Test_splitListToParts(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want []*ListNode
	}{
		{
			name: "Example 1",
			args: args{
				head: listnode.NewG(1, 2, 3),
				k:    5,
			},
			want: []*ListNode{
				listnode.NewG(1),
				listnode.NewG(2),
				listnode.NewG(3),
				nil,
				nil,
			},
		},
		{
			name: "Examle 2",
			args: args{
				head: listnode.NewG(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
				k:    3,
			},
			want: []*ListNode{
				listnode.NewG(1, 2, 3, 4),
				listnode.NewG(5, 6, 7),
				listnode.NewG(8, 9, 10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitListToParts(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitListToParts() = %v, want %v", got, tt.want)
			}
		})
	}
}
