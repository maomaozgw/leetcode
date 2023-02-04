package p1779

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_nearestValidPoint(t *testing.T) {
	type args struct {
		x      int
		y      int
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				x:      3,
				y:      4,
				points: tools.NewGridFromStr[int](strconv.Atoi, `[1,2],[3,1],[2,4],[2,3],[4,4]`),
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				x:      3,
				y:      4,
				points: tools.NewGridFromStr[int](strconv.Atoi, `[3,4]`),
			},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{
				x: 3,
				y: 4,
				points: [][]int{
					{2, 3},
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nearestValidPoint(tt.args.x, tt.args.y, tt.args.points); got != tt.want {
				t.Errorf("nearestValidPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
