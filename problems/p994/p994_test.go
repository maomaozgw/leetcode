package p994

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_orangesRotting(t *testing.T) {
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
				grid: tools.NewGridFromStr(strconv.Atoi, `[2,1,1],[1,1,0],[0,1,1]`),
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				grid: tools.NewGridFromStr(strconv.Atoi, `[2,1,1],[0,1,1],[1,0,1]`),
			},
			want: -1,
		},
		{
			name: "Example 3",
			args: args{
				grid: tools.NewGridFromStr(strconv.Atoi, `[0,2]`),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orangesRotting(tt.args.grid); got != tt.want {
				t.Errorf("orangesRotting() = %v, want %v", got, tt.want)
			}
		})
	}
}
