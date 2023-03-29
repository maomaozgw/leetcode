package p64

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				grid: tools.NewGridFromStr[int](strconv.Atoi, `[1,3,1],[1,5,1],[4,2,1]`),
			},
			want: 7,
		},
		{
			name: "Example 2",
			args: args{
				grid: tools.NewGridFromStr[int](strconv.Atoi, `[1,2,3],[4,5,6]`),
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
