package p840

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_numMagicSquaresInside(t *testing.T) {
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
				grid: tools.NewGridFromStr[int](strconv.Atoi, `[[4,3,8,4],[9,5,1,9],[2,7,6,2]]`),
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				grid: [][]int{
					{8},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numMagicSquaresInside(tt.args.grid); got != tt.want {
				t.Errorf("numMagicSquaresInside() = %v, want %v", got, tt.want)
			}
		})
	}
}
