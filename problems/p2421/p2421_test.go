package p2421

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_numberOfGoodPaths(t *testing.T) {
	type args struct {
		vals  []int
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
				vals:  []int{1, 3, 2, 1, 3},
				edges: tools.NewGridFromStr(strconv.Atoi, `[0,1],[0,2],[2,3],[2,4]`),
			},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{
				vals:  []int{1, 1, 2, 2, 3},
				edges: tools.NewGridFromStr(strconv.Atoi, `[0,1],[1,2],[2,3],[2,4]`),
			},
			want: 7,
		},
		{
			name: "Example 3",
			args: args{
				vals:  []int{1},
				edges: nil,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfGoodPaths(tt.args.vals, tt.args.edges); got != tt.want {
				t.Errorf("numberOfGoodPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
