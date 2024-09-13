package p1310

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_xorQueries(t *testing.T) {
	type args struct {
		arr     []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				arr:     []int{1, 3, 4, 8},
				queries: tools.NewGridFromStr[int](strconv.Atoi, `[0,1],[1,2],[0,3],[3,3]`),
			},
			want: []int{2, 7, 14, 8},
		},
		{
			name: "Example ",
			args: args{
				arr:     []int{4, 8, 2, 10},
				queries: [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}},
			},
			want: []int{8, 0, 4, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorQueries(tt.args.arr, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("xorQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
