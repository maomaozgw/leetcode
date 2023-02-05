package p1232

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_checkStraightLine(t *testing.T) {
	type args struct {
		coordinates [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				coordinates: tools.NewGridFromStr[int](strconv.Atoi, `[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]`),
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				coordinates: tools.NewGridFromStr[int](strconv.Atoi, `[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]`),
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				coordinates: tools.NewGridFromStr[int](strconv.Atoi, `[1,2],[3,4],[2,3],[4,5],[5,6],[6,7]`),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkStraightLine(tt.args.coordinates); got != tt.want {
				t.Errorf("checkStraightLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
