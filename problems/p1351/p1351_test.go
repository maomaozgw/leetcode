package p1351

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_countNegatives(t *testing.T) {
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
				grid: tools.NewGridFromStr[int](strconv.Atoi, `[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]`),
			},
			want: 8,
		},
		{
			name: "Example 2",
			args: args{
				grid: tools.NewGridFromStr[int](strconv.Atoi, `[3,2],[1,0]`),
			},
			want: 0,
		},
		{
			name: "UE 1",
			args: args{
				grid: tools.NewGridFromStr[int](strconv.Atoi, `[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3],[-2,-3,-4,-5]`),
			},
			want: 12,
		},
		{
			name: "UE 2",
			args: args{
				grid: tools.NewGridFromStr[int](strconv.Atoi, `[-1,-1,-2,-3],[-2,-3,-4,-5]`),
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNegatives(tt.args.grid); got != tt.want {
				t.Errorf("countNegatives() = %v, want %v", got, tt.want)
			}
		})
	}
}
