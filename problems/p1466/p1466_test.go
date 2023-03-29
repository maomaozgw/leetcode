package p1466

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_minReorder(t *testing.T) {
	type args struct {
		n           int
		connections [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				n:           6,
				connections: tools.NewGridFromStr[int](strconv.Atoi, `[0,1],[1,3],[2,3],[4,0],[4,5]`),
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				n:           5,
				connections: tools.NewGridFromStr[int](strconv.Atoi, `[1,0],[1,2],[3,2],[3,4]`),
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				n:           3,
				connections: tools.NewGridFromStr[int](strconv.Atoi, `[1,0],[2,0]`),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minReorder(tt.args.n, tt.args.connections); got != tt.want {
				t.Errorf("minReorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
