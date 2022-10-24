package p1293

import (
	"github.com/maomaozgw/leetcode/structures/tools"
	"testing"
)

func Test_shortestPath(t *testing.T) {
	type args struct {
		grid [][]int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				grid: tools.NewGridFromStr("[0,0,0],[1,1,0],[0,0,0],[0,1,1],[0,0,0]"),
				k:    1,
			},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{
				grid: tools.NewGridFromStr("[0,1,1],[1,1,1],[1,0,0]"),
				k:    1,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPath(tt.args.grid, tt.args.k); got != tt.want {
				t.Errorf("shortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
