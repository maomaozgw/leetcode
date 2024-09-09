package p2326

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/listnode"
	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_spiralMatrix(t *testing.T) {
	type args struct {
		m    int
		n    int
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				m:    3,
				n:    5,
				head: listnode.NewG[int](3, 0, 2, 6, 8, 1, 7, 9, 4, 2, 5, 5, 0),
			},
			want: tools.NewGridFromStr[int](strconv.Atoi, `[3,0,2,6,8],[5,0,-1,-1,1],[5,2,4,9,7]`),
		},
		{
			name: "Example 2",
			args: args{
				m:    1,
				n:    4,
				head: listnode.NewG(0, 1, 2),
			},
			want: [][]int{
				{0, 1, 2, -1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralMatrix(tt.args.m, tt.args.n, tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
