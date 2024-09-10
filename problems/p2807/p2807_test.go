package p2807

import (
	"reflect"
	"testing"

	"github.com/maomaozgw/leetcode/structures/listnode"
)

func Test_insertGreatestCommonDivisors(t *testing.T) {
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
				head: listnode.NewG(18, 6, 10, 3),
			},
			want: listnode.NewG(18, 6, 6, 2, 10, 1, 3),
		},
		{
			name: "Example 2",
			args: args{
				head: listnode.NewG(7),
			},
			want: listnode.NewG(7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertGreatestCommonDivisors(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertGreatestCommonDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
