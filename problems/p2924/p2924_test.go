package p2924

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_findChampion(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				n:     3,
				edges: tools.NewGridFromStr[int](strconv.Atoi, `[0,1],[1,2]`),
			},
			want: 0,
		},
		{
			name: "Example 2",
			args: args{
				n:     4,
				edges: tools.NewGridFromStr[int](strconv.Atoi, `[0,2],[1,3],[1,2]`),
			},
			want: -1,
		},
		{
			name: "WA1",
			args: args{
				n:     1,
				edges: [][]int{},
			},
			want: 0,
		},
		{
			name: "WA2",
			args: args{
				n:     3,
				edges: tools.NewGridFromStr[int](strconv.Atoi, `[[0,1]]`),
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findChampion(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("findChampion() = %v, want %v", got, tt.want)
			}
		})
	}
}
