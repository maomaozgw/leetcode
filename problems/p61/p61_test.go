package p61

import (
	"reflect"
	"testing"

	"github.com/maomaozgw/leetcode/structures/listnode"
)

func Test_rotateRight(t *testing.T) {
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
			want: listnode.NewG[int](4, 5, 1, 2, 3),
		},
		{
			name: "Example 2",
			args: args{
				head: listnode.NewG(0, 1, 2),
				k:    4,
			},
			want: listnode.NewG(2, 0, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
