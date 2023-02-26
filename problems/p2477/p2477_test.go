package p2477

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_minimumFuelCost(t *testing.T) {
	type args struct {
		roads [][]int
		seats int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{
				roads: tools.NewGridFromStr[int](strconv.Atoi, `[0,1],[0,2],[0,3]`),
				seats: 5,
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				roads: tools.NewGridFromStr[int](strconv.Atoi, `[3,1],[3,2],[1,0],[0,4],[0,5],[4,6]`),
				seats: 2,
			},
			want: 7,
		},
		{
			name: "Example 3",
			args: args{
				roads: [][]int{},
				seats: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumFuelCost(tt.args.roads, tt.args.seats); got != tt.want {
				t.Errorf("minimumFuelCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
