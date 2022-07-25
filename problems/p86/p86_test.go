package p86

import (
	"code.byted.org/zhaoguowei/leetcode/structures/listnode"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example",
			args: args{
				head: listnode.NewG(1, 4, 3, 2, 5, 2),
				x:    3,
			},
			want: listnode.NewG(1, 2, 2, 4, 3, 5),
		},
		{
			name: "no smaller",
			args: args{
				head: listnode.NewG(5, 4, 3, 3, 5, 3),
				x:    3,
			},
			want: listnode.NewG(5, 4, 3, 3, 5, 3),
		},
		{
			name: "no bigger",
			args: args{
				head: listnode.NewG(5, 4, 3, 3, 5, 3),
				x:    10,
			},
			want: listnode.NewG(5, 4, 3, 3, 5, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !listnode.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
