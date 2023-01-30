package p148

import (
	"reflect"
	"testing"

	"github.com/maomaozgw/leetcode/structures/listnode"
)

func Test_sortList(t *testing.T) {
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
				head: listnode.NewG(4, 2, 1, 3),
			},
			want: listnode.NewG(1, 2, 3, 4),
		},
		{
			name: "Example 2",
			args: args{
				head: listnode.NewG(-1, 5, 3, 4, 0),
			},
			want: listnode.NewG(-1, 0, 3, 4, 5),
		},
		{
			name: "Example 3",
			args: args{
				head: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
