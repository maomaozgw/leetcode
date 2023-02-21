package p815

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_numBusesToDestination(t *testing.T) {
	type args struct {
		routes [][]int
		source int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				routes: tools.NewGridFromStr[int](strconv.Atoi, `[1,2,7],[3,6,7]`),
				source: 1,
				target: 6,
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				routes: tools.NewGridFromStr[int](strconv.Atoi, `[7,12],[4,5,15],[6],[15,19],[9,12,13]`),
				source: 15,
				target: 12,
			},
			want: -1,
		},
		{
			name: "WA 1",
			args: args{
				routes: tools.NewGridFromStr[int](strconv.Atoi, `[1,7],[3,5]`),
				source: 5,
				target: 5,
			},
			want: 0,
		},
		{
			name: "UE 1",
			args: args{
				routes: tools.NewGridFromStr[int](strconv.Atoi, `[1,7],[3,5]`),
				source: 5,
				target: 3,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numBusesToDestination(tt.args.routes, tt.args.source, tt.args.target); got != tt.want {
				t.Errorf("numBusesToDestination() = %v, want %v", got, tt.want)
			}
		})
	}
}
