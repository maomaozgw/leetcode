package p1162

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_maxDistance(t *testing.T) {
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
				grid: tools.NewGridFromStr[int](strconv.Atoi, `[1,0,1],[0,0,0],[1,0,1]`),
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				grid: tools.NewGridFromStr[int](strconv.Atoi, `[1,0,0],[0,0,0],[0,0,0]`),
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.args.grid); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
