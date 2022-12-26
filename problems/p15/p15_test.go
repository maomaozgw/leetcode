package p15

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/maomaozgw/leetcode/structures/tools"
	"strconv"
	"testing"
)

func Test_threeSum(t *testing.T) {
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
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: tools.NewGridFromStr(strconv.Atoi, `[-1,-1,2],[-1,0,1]`),
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{0, 1, 1},
			},
			want: [][]int{},
		},
		{
			name: "Example 3",
			args: args{
				[]int{0, 0, 0},
			},
			want: [][]int{{0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !cmp.Equal(got, tt.want, cmpopts.EquateEmpty()) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
