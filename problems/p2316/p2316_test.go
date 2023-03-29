package p2316

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_countPairs(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{
				n:     3,
				edges: tools.NewGridFromStr[int](strconv.Atoi, `[0,1],[0,2],[1,2]`),
			},
			want: 0,
		},
		{
			name: "Example 2",
			args: args{
				n:     7,
				edges: tools.NewGridFromStr[int](strconv.Atoi, `[0,2],[0,5],[2,4],[1,6],[5,4]`),
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
