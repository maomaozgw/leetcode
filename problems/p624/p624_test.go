package p624

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_maxDistance(t *testing.T) {
	type args struct {
		arrays [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				arrays: tools.NewGridFromStr[int](strconv.Atoi, `[1,2,3],[4,5],[1,2,3]`),
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				arrays: tools.NewGridFromStr[int](strconv.Atoi, `[1],[1]`),
			},
			want: 0,
		},
		{
			name: "WA 1",
			args: args{
				arrays: tools.NewGridFromStr[int](strconv.Atoi, `[3,4],[1,5]`),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.args.arrays); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
