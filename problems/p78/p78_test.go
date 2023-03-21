package p78

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: tools.NewGridFromStr[int](strconv.Atoi, `[],[1],[2],[3],[1,2],[1,3],[2,3],[1,2,3]`),
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{0},
			},
			want: [][]int{{}, {0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.nums); !cmp.Equal(got, tt.want, cmpopts.EquateEmpty()) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
