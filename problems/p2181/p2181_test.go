package p2181

import (
	"reflect"
	"testing"

	"github.com/maomaozgw/leetcode/structures/listnode"
)

func Test_mergeNodes(t *testing.T) {
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
				head: listnode.NewG[int](0, 3, 1, 0, 4, 5, 2, 0),
			},
			want: listnode.NewG[int](4, 11),
		},
		{
			name: "Example 2",
			args: args{
				head: listnode.NewG(0, 1, 0, 3, 0, 2, 2, 0),
			},
			want: listnode.NewG(1, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
