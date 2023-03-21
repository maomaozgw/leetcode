package p1091

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_shortestPathBinaryMatrix(t *testing.T) {
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
				grid: tools.NewGridFromStr[int](strconv.Atoi, `[0,1],[1,0]`),
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				grid: tools.NewGridFromStr[int](strconv.Atoi, `[0,0,0],[1,1,0],[1,1,0]`),
			},
			want: 4,
		},
		{
			name: "Example 3",
			args: args{
				grid: tools.NewGridFromStr[int](strconv.Atoi, `[1,0,0],[1,1,0],[1,1,0]`),
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPathBinaryMatrix(tt.args.grid); got != tt.want {
				t.Errorf("shortestPathBinaryMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
