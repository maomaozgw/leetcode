package p1975

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_maxMatrixSum(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{
				matrix: tools.NewGridFromStr[int](strconv.Atoi, `[1,-1],[-1,1]`),
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				matrix: tools.NewGridFromStr[int](strconv.Atoi, `[1,2,3],[-1,-2,-3],[1,2,3]`),
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxMatrixSum(tt.args.matrix); got != tt.want {
				t.Errorf("maxMatrixSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
