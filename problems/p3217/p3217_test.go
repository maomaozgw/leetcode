package p3217

import (
	"reflect"
	"testing"

	"github.com/maomaozgw/leetcode/structures/listnode"
)

func Test_modifiedList(t *testing.T) {
	type args struct {
		nums []int
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1 ",
			args: args{
				nums: []int{
					1, 2, 3,
				},
				head: listnode.NewG(1, 2, 3, 4, 5),
			},
			want: listnode.NewG(4, 5),
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{1},
				head: listnode.NewG(1, 2, 1, 2, 1, 2),
			},
			want: listnode.NewG(2, 2, 2),
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{5},
				head: listnode.NewG(1, 2, 3, 4),
			},
			want: listnode.NewG(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modifiedList(tt.args.nums, tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("modifiedList() = %v, want %v", got, tt.want)
			}
		})
	}
}
