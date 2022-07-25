package p92

import (
	"testing"

	"github.com/maomaozgw/leetcode/structures/listnode"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example",
			args: args{
				head:  listnode.NewG(1, 2, 3, 4, 5),
				left:  2,
				right: 4,
			},
			want: listnode.NewG(1, 4, 3, 2, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !listnode.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
