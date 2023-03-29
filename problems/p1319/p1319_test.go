package p1319

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_makeConnected(t *testing.T) {
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
				n:           4,
				connections: tools.NewGridFromStr[int](strconv.Atoi, `[0,1],[0,2],[1,2]`),
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				n:           6,
				connections: tools.NewGridFromStr[int](strconv.Atoi, `[0,1],[0,2],[0,3],[1,2],[1,3]`),
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				n:           6,
				connections: tools.NewGridFromStr[int](strconv.Atoi, `[0,1],[0,2],[0,3],[1,2]`),
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeConnected(tt.args.n, tt.args.connections); got != tt.want {
				t.Errorf("makeConnected() = %v, want %v", got, tt.want)
			}
		})
	}
}
